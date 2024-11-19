package entity

import (
	"bnqkl/chain-cms/database/model"
	"bnqkl/chain-cms/exception"
	"bnqkl/chain-cms/helper"
	"bnqkl/chain-cms/logger"
	"bnqkl/chain-cms/redis"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type EntityService struct {
	db  *gorm.DB
	log *logger.Logger
}

var entityService *EntityService

func NewEntityService(db *gorm.DB, log *logger.Logger) {
	entityService = &EntityService{
		db:  db,
		log: log,
	}
}

func GetEntityService() *EntityService {
	return entityService
}

func (s *EntityService) getFactoryRedisLock(chainName, chainMagic, factoryId string) string {
	return fmt.Sprintf("lock:factory:%s_%s_%s", chainName, chainMagic, factoryId)
}

func (s *EntityService) getEntityRedisLock(chainName, chainMagic, factoryId, entityId string) string {
	return fmt.Sprintf("lock:entity:%s_%s_%s_%s", chainName, chainMagic, factoryId, entityId)
}

func (s *EntityService) Add(req AddEntityReq) (AddEntityRes, error) {
	chainName := req.GetChainName()
	chainMagic := req.GetChainMagic()
	factoryId := req.GetFactoryId()
	entityId := req.GetEntityId()
	return redis.DoWithLockMulti([]string{s.getFactoryRedisLock(chainName, chainMagic, factoryId), s.getEntityRedisLock(chainName, chainMagic, factoryId, entityId)}, func() (AddEntityRes, error) {
		var entity model.Entity
		var count int64
		result := s.db.Model(&entity).Where(map[string]any{
			entity.GetChainNameColumnName():  chainName,
			entity.GetChainMagicColumnName(): chainMagic,
			entity.GetFactoryIdColumnName():  factoryId,
			entity.GetEntityIdColumnName():   entityId,
		}).Count(&count)
		if result.Error != nil {
			s.log.Error(result.Error)
			return false, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_QUERY_FAILED)
		}
		entityFlag := fmt.Sprintf("%s-%s-%s-%s", chainName, chainMagic, factoryId, entityId)
		if count > 0 {
			return false, exception.NewExceptionWithParam(exception.THE_ENTITY_IS_ALREADY_EXISTED, map[string]string{
				"entity": entityFlag,
			})
		}
		entity = model.NewEntity(req.GetPossessor(), chainName, chainMagic, req.GetFactoryName(), factoryId, entityId, req.GetTaxCollector(), req.GetTaxAssetPrealnum(), req.GetType(), req.GetHash(), req.GetExtension())
		result = s.db.Create(&entity)
		if result.Error != nil {
			s.log.Error(result.Error)
			return false, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_ADD_FAILED)
		}
		if result.RowsAffected == 0 {
			return false, exception.NewExceptionWithParam(exception.THE_ENTITY_ADD_FAILED, map[string]string{
				"entity": entityFlag,
			})
		}
		return true, nil
	})
}
func (s *EntityService) AddMulti(req AddEntityMultiReq) (AddEntityMultiRes, error) {
	chainName := req.GetChainName()
	chainMagic := req.GetChainMagic()
	factoryId := req.GetFactoryId()
	return redis.DoWithLock(s.getFactoryRedisLock(chainName, chainMagic, factoryId), func() (AddEntityMultiRes, error) {
		var entity model.Entity
		entities := req.GetEntities()
		entityIds := []string{}
		for _, item := range entities {
			entityIds = append(entityIds, item.GetEntityId())
		}
		entityFlag := fmt.Sprintf("%s-%s-%s", chainName, chainMagic, factoryId)
		err := helper.SafetyRunTask(entityIds, 100, func(chunkDatas []string) error {
			var count int64
			result := s.db.Model(&entity).Where(map[string]any{
				entity.GetChainNameColumnName():  chainName,
				entity.GetChainMagicColumnName(): chainMagic,
				entity.GetFactoryIdColumnName():  factoryId,
				entity.GetEntityIdColumnName():   chunkDatas,
			}).Count(&count)
			if result.Error != nil {
				s.log.Error(result.Error)
				return exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_QUERY_FAILED)
			}
			if count > 0 {
				return exception.NewExceptionWithParam(exception.SOME_ENTITY_IS_ALREADY_EXISTED, map[string]string{
					"entity": entityFlag,
				})
			}
			return nil
		})
		if err != nil {
			return false, err
		}
		possessor := req.GetPossessor()
		factoryName := req.GetFactoryName()
		taxCollector := req.GetTaxCollector()
		entityType := req.GetType()
		hash := req.GetHash()
		extension := req.GetExtension()
		_entities := []model.Entity{}
		for _, item := range entities {
			_taxCollector := taxCollector
			if item.TaxCollector != nil {
				_taxCollector = item.GetTaxCollector()
			}
			entity = model.NewEntity(possessor, chainName, chainMagic, factoryName, factoryId, item.GetEntityId(), _taxCollector, item.GetTaxAssetPrealnum(), entityType, hash, extension)
			_entities = append(_entities, entity)
		}
		err = s.db.Transaction(func(tx *gorm.DB) error {
			return helper.SafetyRunTask(_entities, 1000, func(chunkDatas []model.Entity) error {
				result := tx.Create(&chunkDatas)
				if result.Error != nil {
					s.log.Error(result.Error)
					return exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_ADD_FAILED)
				}
				if result.RowsAffected == 0 {
					return exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_ADD_FAILED)
				}
				return nil
			})
		})
		if err != nil {
			return false, err
		}
		return true, nil
	})
}

func (s *EntityService) Update(req UpdateEntityReq) (UpdateEntityRes, error) {
	entity := model.Entity{}
	hasParams := false
	condition := map[string]any{}
	if req.Possessor != nil {
		hasParams = true
		condition[entity.GetPossessorColumnName()] = req.GetPossessor()
	}
	if req.FactoryName != nil {
		hasParams = true
		condition[entity.GetFactoryNameColumnName()] = req.GetFactoryName()
	}
	if !hasParams {
		return false, exception.NewExceptionWithoutParam(exception.NO_UPDATE_CONDITIONS)
	}
	chainName := req.GetChainName()
	chainMagic := req.GetChainMagic()
	factoryId := req.GetFactoryId()
	entityId := req.GetEntityId()
	return redis.DoWithLock(s.getEntityRedisLock(chainName, chainMagic, factoryId, entityId), func() (UpdateEntityRes, error) {
		var count int64
		filter := map[string]any{
			entity.GetChainNameColumnName():  chainName,
			entity.GetChainMagicColumnName(): chainMagic,
			entity.GetFactoryIdColumnName():  factoryId,
			entity.GetEntityIdColumnName():   entityId,
		}
		result := s.db.Model(&entity).Where(filter).Count(&count)
		if result.Error != nil {
			s.log.Error(result.Error)
			return false, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_QUERY_FAILED)
		}
		entityFlag := fmt.Sprintf("%s-%s-%s-%s", chainName, chainMagic, factoryId, entityId)
		if count <= 0 {
			return false, exception.NewExceptionWithParam(exception.THE_ENTITY_IS_NOT_EXISTED, map[string]string{
				"entity": entityFlag,
			})
		}
		result = s.db.Model(&entity).Where(filter).Updates(condition)
		if result.Error != nil {
			s.log.Error(result.Error)
			return false, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_UPDATE_FAILED)
		}
		if result.RowsAffected == 0 {
			return false, exception.NewExceptionWithParam(exception.THE_ENTITY_UPDATE_FAILED, map[string]string{
				"entity": entityFlag,
			})
		}
		return true, nil
	})
}

func (s *EntityService) GetUserFactoryAll(req GetUserFactoryAllReq) (GetUserFactoryAllRes, error) {
	entity := model.Entity{}
	possessor := req.GetPossessor()
	entities := []model.Entity{}
	res := GetUserFactoryAllRes{
		Factories: []UserFactoryInfo{},
	}
	result := s.db.Model(&entity).Where(map[string]any{
		entity.GetPossessorColumnName(): possessor,
	}).Find(&entities)
	if result.Error != nil {
		s.log.Error(result.Error)
		return res, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_QUERY_FAILED)
	}
	if result.RowsAffected == 0 {
		return res, nil
	}
	factoryMap := map[string]*UserFactoryInfo{}
	getMapKey := func(entity model.Entity) string {
		var mapKey strings.Builder
		mapKey.WriteString(entity.ChainName)
		mapKey.WriteString("_")
		mapKey.WriteString(entity.ChainMagic)
		mapKey.WriteString("_")
		mapKey.WriteString(entity.FactoryId)
		return mapKey.String()
	}
	for _, entity := range entities {
		mapKey := getMapKey(entity)
		factory, ok := factoryMap[mapKey]
		if ok {
			factory.NumberOfEntities += 1
		} else {
			factory = &UserFactoryInfo{
				ChainName:        entity.ChainName,
				ChainMagic:       entity.ChainMagic,
				FactoryId:        entity.FactoryId,
				FactoryName:      entity.FactoryName,
				NumberOfEntities: 1,
			}
			factoryMap[mapKey] = factory
		}
	}
	factories := []UserFactoryInfo{}
	for _, factory := range factoryMap {
		factories = append(factories, *factory)
	}
	res.Factories = factories
	return res, nil
}

func (s *EntityService) GetUserFactoryEntityAll(req GetUserFactoryEntityAllReq) (GetUserFactoryEntityAllRes, error) {
	entity := model.Entity{}
	possessor := req.GetPossessor()
	entities := []model.Entity{}
	res := GetUserFactoryEntityAllRes{
		Entities: []UserFactoryEntityInfo{},
	}
	condition := map[string]any{
		entity.GetPossessorColumnName():  possessor,
		entity.GetChainNameColumnName():  req.GetChainName(),
		entity.GetChainMagicColumnName(): req.GetChainMagic(),
	}
	if req.FactoryId != nil {
		condition[entity.GetFactoryIdColumnName()] = req.GetFactoryId()
	}
	if req.Type != nil {
		condition[entity.GetTypeColumnName()] = req.GetType()
	}
	result := s.db.Model(&entity).Where(condition).Find(&entities)
	if result.Error != nil {
		s.log.Error(result.Error)
		return res, exception.NewExceptionWithoutParam(exception.SYSTEM_DATABASE_QUERY_FAILED)
	}
	if result.RowsAffected == 0 {
		return res, nil
	}
	// shit 限量集合并
	limitMap := map[string][]SubEntity{}
	factoryMap := map[string]model.Entity{}
	userEntities := []UserFactoryEntityInfo{}
	for _, entity := range entities {
		// 非限量集
		if entity.Type == 1 {
			userEntities = append(userEntities, NewUserFactoryEntityInfo(entity))
			continue
		}
		// 限量集
		factoryId := entity.FactoryId
		subEntities, ok := limitMap[factoryId]
		subEntity := SubEntity{
			EntityId:         entity.EntityId,
			TaxCollector:     entity.TaxCollector,
			TaxAssetPrealnum: entity.TaxAssetPrealnum,
		}
		if ok {
			subEntities = append(subEntities, subEntity)
		} else {
			subEntities = []SubEntity{subEntity}
			factoryMap[factoryId] = entity
		}
		limitMap[factoryId] = subEntities
	}
	for factoryId, subEntities := range limitMap {
		entity := factoryMap[factoryId]
		userEntities = append(userEntities, UserFactoryEntityInfo{
			Possessor:        entity.Possessor,
			ChainName:        entity.ChainName,
			ChainMagic:       entity.ChainMagic,
			FactoryName:      entity.FactoryName,
			FactoryId:        entity.FactoryId,
			EntityId:         "",
			TaxCollector:     "",
			TaxAssetPrealnum: "0",
			Type:             entity.Type,
			Hash:             entity.Hash,
			Extension:        entity.Extension,
			SubEntities:      subEntities,
		})
	}
	res.Entities = userEntities
	return res, nil
}
