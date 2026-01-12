package main

import "fmt"

func (a *analytics) handleEmailBounce(em email) error {
	value := em.recipient.updateStatus(em.status)
	if value != nil {
		return fmt.Errorf("error updating user status: %w", value)
	}
	track := a.track(em.status)
	if track != nil {
		return fmt.Errorf("error tracking user bounce: %w", track)
	}
	return nil
}
