// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"errors"
	"fmt"
	"github.com/dubinc/dub-go/internal/utils"
)

type WebhookEventType string

const (
	WebhookEventTypeLinkWebhookEvent    WebhookEventType = "LinkWebhookEvent"
	WebhookEventTypeLinkClickedEvent    WebhookEventType = "LinkClickedEvent"
	WebhookEventTypeLeadCreatedEvent    WebhookEventType = "LeadCreatedEvent"
	WebhookEventTypeSaleCreatedEvent    WebhookEventType = "SaleCreatedEvent"
	WebhookEventTypePartnerCreatedEvent WebhookEventType = "PartnerCreatedEvent"
)

// WebhookEvent - Webhook event schema
type WebhookEvent struct {
	LinkWebhookEvent    *LinkWebhookEvent    `queryParam:"inline"`
	LinkClickedEvent    *LinkClickedEvent    `queryParam:"inline"`
	LeadCreatedEvent    *LeadCreatedEvent    `queryParam:"inline"`
	SaleCreatedEvent    *SaleCreatedEvent    `queryParam:"inline"`
	PartnerCreatedEvent *PartnerCreatedEvent `queryParam:"inline"`

	Type WebhookEventType
}

func CreateWebhookEventLinkWebhookEvent(linkWebhookEvent LinkWebhookEvent) WebhookEvent {
	typ := WebhookEventTypeLinkWebhookEvent

	return WebhookEvent{
		LinkWebhookEvent: &linkWebhookEvent,
		Type:             typ,
	}
}

func CreateWebhookEventLinkClickedEvent(linkClickedEvent LinkClickedEvent) WebhookEvent {
	typ := WebhookEventTypeLinkClickedEvent

	return WebhookEvent{
		LinkClickedEvent: &linkClickedEvent,
		Type:             typ,
	}
}

func CreateWebhookEventLeadCreatedEvent(leadCreatedEvent LeadCreatedEvent) WebhookEvent {
	typ := WebhookEventTypeLeadCreatedEvent

	return WebhookEvent{
		LeadCreatedEvent: &leadCreatedEvent,
		Type:             typ,
	}
}

func CreateWebhookEventSaleCreatedEvent(saleCreatedEvent SaleCreatedEvent) WebhookEvent {
	typ := WebhookEventTypeSaleCreatedEvent

	return WebhookEvent{
		SaleCreatedEvent: &saleCreatedEvent,
		Type:             typ,
	}
}

func CreateWebhookEventPartnerCreatedEvent(partnerCreatedEvent PartnerCreatedEvent) WebhookEvent {
	typ := WebhookEventTypePartnerCreatedEvent

	return WebhookEvent{
		PartnerCreatedEvent: &partnerCreatedEvent,
		Type:                typ,
	}
}

func (u *WebhookEvent) UnmarshalJSON(data []byte) error {

	var linkWebhookEvent LinkWebhookEvent = LinkWebhookEvent{}
	if err := utils.UnmarshalJSON(data, &linkWebhookEvent, "", true, false); err == nil {
		u.LinkWebhookEvent = &linkWebhookEvent
		u.Type = WebhookEventTypeLinkWebhookEvent
		return nil
	}

	var linkClickedEvent LinkClickedEvent = LinkClickedEvent{}
	if err := utils.UnmarshalJSON(data, &linkClickedEvent, "", true, false); err == nil {
		u.LinkClickedEvent = &linkClickedEvent
		u.Type = WebhookEventTypeLinkClickedEvent
		return nil
	}

	var leadCreatedEvent LeadCreatedEvent = LeadCreatedEvent{}
	if err := utils.UnmarshalJSON(data, &leadCreatedEvent, "", true, false); err == nil {
		u.LeadCreatedEvent = &leadCreatedEvent
		u.Type = WebhookEventTypeLeadCreatedEvent
		return nil
	}

	var saleCreatedEvent SaleCreatedEvent = SaleCreatedEvent{}
	if err := utils.UnmarshalJSON(data, &saleCreatedEvent, "", true, false); err == nil {
		u.SaleCreatedEvent = &saleCreatedEvent
		u.Type = WebhookEventTypeSaleCreatedEvent
		return nil
	}

	var partnerCreatedEvent PartnerCreatedEvent = PartnerCreatedEvent{}
	if err := utils.UnmarshalJSON(data, &partnerCreatedEvent, "", true, false); err == nil {
		u.PartnerCreatedEvent = &partnerCreatedEvent
		u.Type = WebhookEventTypePartnerCreatedEvent
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for WebhookEvent", string(data))
}

func (u WebhookEvent) MarshalJSON() ([]byte, error) {
	if u.LinkWebhookEvent != nil {
		return utils.MarshalJSON(u.LinkWebhookEvent, "", true)
	}

	if u.LinkClickedEvent != nil {
		return utils.MarshalJSON(u.LinkClickedEvent, "", true)
	}

	if u.LeadCreatedEvent != nil {
		return utils.MarshalJSON(u.LeadCreatedEvent, "", true)
	}

	if u.SaleCreatedEvent != nil {
		return utils.MarshalJSON(u.SaleCreatedEvent, "", true)
	}

	if u.PartnerCreatedEvent != nil {
		return utils.MarshalJSON(u.PartnerCreatedEvent, "", true)
	}

	return nil, errors.New("could not marshal union type WebhookEvent: all fields are null")
}
