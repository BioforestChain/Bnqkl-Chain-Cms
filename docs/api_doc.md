# api 接口说明

## entity 模块

### 新增 entity

```ts
/**
 * 新增 entity
 *
 * method: post
 *
 * path: /api/entity/add
 */
export interface AddEntityReq {
  /**拥有者，min=1，max=255 */
  possessor: string;
  /**链名，min=1，max=30 */
  chainName: string;
  /**链网络标识，min=1，max=30 */
  chainMagic: string;
  /**模板名称，min=1，max=50 */
  factoryName?: string;
  /**模板ID，min=1，max=50 */
  factoryId: string;
  /**非同质资产ID，min=1，max=50 */
  entityId: string;
  /**收税人，min=1，max=255 */
  taxCollector: string;
  /**缴纳数量，min=1，max=50 */
  taxAssetPrealnum: string;
  /**类型(1:普通，2:限量，默认:1) ，min=1，max=2 */
  type: number;
  /**blob哈希，min=1，max=100 */
  hash: string;
  /**blob扩展名，min=1，max=50 */
  extension?: string;
}
export type AddEntityRes = boolean;
```

### 批量新增 entity

```ts
export interface EntityStruct {
  /**非同质资产ID，min=1，max=50 */
  entityId: string;
  /**收税人，min=1，max=255 */
  taxCollector?: string;
  /**缴纳数量，min=1，max=50 */
  taxAssetPrealnum: string;
}
/**
 * 批量新增 entity
 *
 * method: post
 *
 * path: /api/entity/add/multi
 */
export interface AddEntityMultiReq {
  /**拥有者，min=1,max=255 */
  possessor: string;
  /**链名，min=1,max=30 */
  chainName: string;
  /**链网络标识，min=1,max=30 */
  chainMagic: string;
  /**模板名称，min=1,max=50 */
  factoryName?: string;
  /**模板ID，min=1,max=50 */
  factoryId: string;
  /**收税人，min=1,max=255 */
  taxCollector: string;
  /**非同质资产列表，gt=0,lte=6000 */
  entities: EntityStruct[];
  /**类型(1:普通,2:限量,默认:1)，min=1,max=2 */
  type: number;
  /**blob哈希，min=1,max=100 */
  hash: string;
  /**blob扩展名，min=1,max=50 */
  extension: string;
}
export type AddEntityMultiRes = boolean;
```

### 更新 entity

```ts
/**
 * 更新 entity
 *
 * method: post
 *
 * path: /api/entity/update
 */
export interface UpdateEntityReq {
  /**链名，min=1，max=30 */
  chainName: string;
  /**链网络标识，min=1，max=30 */
  chainMagic: string;
  /**模板ID，min=1，max=50 */
  factoryId: string;
  /**非同质资产ID，min=1，max=50 */
  entityId: string;
  /**拥有者，min=1，max=255 */
  possessor?: string;
  /**模板名称，min=1，max=50 */
  factoryName?: string;
}
export type UpdateEntityRes = boolean;
```

### 获取用户持有的 factory 总览

```ts
/**
 * 获取用户持有的 factory 总览
 *
 * method: get
 *
 * path: /api/entity/factory/all
 */
export interface GetUserFactoryAllReq {
  /**拥有者，min=1，max=255 */
  possessor: string;
}
export interface UserFactoryInfo {
  /**链名 */
  chainName: string;
  /**链网络标识 */
  chainMagic: string;
  /**模板ID */
  factoryId: string;
  /**模板名称 */
  factoryName: string;
  /**持有的 entity 数量 */
  numberOfEntities: number;
}
export interface GetUserFactoryAllRes {
  /**factory 列表 */
  factories: UserFactoryInfo[];
}
```

### 获取用户持有的某个 factory 下的所有 entity

```ts
/**
 * 获取用户持有的（某个 factory 下的）所有 entity
 *
 * method: get
 *
 * path: /api/entity/factory/entity/all
 */
export interface GetUserFactoryEntityAllReq {
  /**拥有者，min=1,max=255 */
  possessor: string;
  /**链名，min=1,max=30 */
  chainName: string;
  /**链网络标识，min=1,max=30 */
  chainMagic: string;
  /**模板ID，min=1,max=50 */
  factoryId?: string;
  /**类型(1:普通,2:限量,默认:1)，min=1,max=2 */
  type?: number;
}
export interface SubEntity {
  /**非同质资产ID */
  entityId: string;
  /**收税人 */
  taxCollector: string;
  /**缴纳数量 */
  taxAssetPrealnum: string;
}
export interface UserFactoryEntityInfo {
  /**链名 */
  chainName: string;
  /**链网络标识 */
  chainMagic: string;
  /**模板ID */
  factoryId: string;
  /**模板名称 */
  factoryName: string;
  /**非同质资产ID */
  entityId: string;
  /**拥有者 */
  possessor: string;
  /**收税人 */
  taxCollector: string;
  /**缴纳数量 */
  taxAssetPrealnum: string;
  /**类型(1:普通,2:限量,默认:1) */
  type: number;
  /**blob哈希 */
  hash: string;
  /**blob扩展名 */
  extension: string;
  /**子集 */
  SubEntities: SubEntity[];
}
export interface GetUserFactoryEntityAllRes {
  /**factory/entites 列表 */
  entities: UserFactoryEntityInfo[];
}
```

## 附件模块

### 上传 blob

```ts
/**
 * 上传 blob
 *
 * method: post
 *
 * path: /api/attach/upload/blob
 */
export interface UploadBlobReq {
  /**blob名称 */
  name: string;
  /**blob扩展名 */
  extension?: string;
  /**blob */
  file: any;
}
export interface UploadBlobRes {
  /**文件路径 */
  url: string;
}
```

### 获取 blob

```ts
http://127.0.0.1:3000/blob/name.jpg
```
