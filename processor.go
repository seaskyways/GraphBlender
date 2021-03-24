package GraphBlender

type Processor struct {
	df dfType
}

func NewProcessor(df dfType) *Processor {
	return &Processor{df}
}

type CleanHeadersOptions struct{}

func (p Processor) CleanHeaders(opts ...CleanHeadersOptions) error {
	//row := p.df.Row(0, false)
	//for _, series := range p.df.Series {
	//	series.Rename(fixGraphQLName(series.Name()))
	//}

	return nil
}
