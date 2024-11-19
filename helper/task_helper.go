package helper

func SafetyRunTask[T any](datas []T, limit int, task func(chunkDatas []T) error) error {
	total := len(datas)
	if total == 0 {
		return nil
	}
	start := 0
	end := start + limit
	if end > total {
		end = total
	}
	for {
		chunkDatas := datas[start:end]
		if len(chunkDatas) == 0 {
			break
		}
		err := task(chunkDatas)
		if err != nil {
			return err
		}
		start = end
		end = start + limit
		if end > total {
			end = total
		}
	}
	return nil
}
