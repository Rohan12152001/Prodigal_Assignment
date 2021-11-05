package extract

type TransformManager interface {
	TransformData(data string) error
}