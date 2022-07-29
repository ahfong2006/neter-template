package cron

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

type Cron struct {
	logger *zap.SugaredLogger
	cron   *cron.Cron
}

func (c *Cron) Close() error {
	c.logger.Infof("Closing cron")
	c.cron.Stop()
	return nil
}

func (c *Cron) AddSecondsFunc(sec int, f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@every %ds", sec), f[i])
	}
	return err
}
func (c *Cron) AddMinuteFunc(m int, f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@every %dm", m), f[i])
	}
	return err
}
func (c *Cron) AddHourly(f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@hourly"), f[i])
	}
	return err

}
func (c *Cron) AddDaily(f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@daily"), f[i])
	}
	return err
}
func (c *Cron) AddMonthly(f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@monthly"), f[i])
	}
	return err
}
func (c *Cron) AddYearly(f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@yearly"), f[i])
	}
	return err
}
func (c *Cron) AddWeekly(f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(fmt.Sprintf("@weekly"), f[i])
	}
	return err
}

func (c *Cron) AddManually(expr string, f ...func()) error {
	var err error
	for i := 0; i < len(f); i++ {
		_, err = c.cron.AddFunc(expr, f[i])
	}
	return err
}
func (c *Cron) Run() {
	c.logger.Infof("Starting cron")

	c.cron.Start()
}
func NewCron(logger *zap.SugaredLogger) *Cron {
	c := &Cron{logger: logger, cron: cron.New(
		cron.WithParser(cron.NewParser(
			cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
		)), cron.WithChain(cron.Recover(cron.DefaultLogger)),
	)}
	return c
}
