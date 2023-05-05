package processor

import (
	"context"
	"errors"
	"fmt"
)

func (p *Processor) BlacklistCheck(ctx context.Context) error {
	caller := p.agi.Env["callerid"]
	if caller == "" {
		return errors.New("callerID is not defined")
	}

	blacklisted, err := p.storage.BlacklistCheck(ctx, caller)
	if err != nil {
		return err
	}

	if err = p.Verbose(fmt.Sprintf("My blacklist is %v", blacklisted)); err != nil {
		return err
	}

	if blacklisted == 1 {

	}
	return nil
}
