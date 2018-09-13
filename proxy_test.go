package goka

type nullProxy struct{}

func (p *nullProxy) Add(topic string, offset int64) error     { return nil }
func (p *nullProxy) Remove(topic string) error                { return nil }
func (p *nullProxy) LowWaterMark(topic string) (int64, error) { return 0, nil }
func (p *nullProxy) AddGroup()                                {}
func (p *nullProxy) Stop()                                    {}
