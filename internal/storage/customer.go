package storage

import (
	"github.com/albertowusuasare/customer-app/internal/adding"
	"github.com/albertowusuasare/customer-app/internal/retrieving"
	"github.com/albertowusuasare/customer-app/internal/updating"
	"github.com/albertowusuasare/customer-app/internal/uuid"
)

type InsertCustomerFunc func(unPersistedCustomer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer

type RetrieveCustomerFunc func(customerId string) retrieving.Customer

type RetrieveCustomersFunc func(request retrieving.MultiRequest) []retrieving.Customer

type UpdateCustomerFunc func(request updating.Request) updating.UpdatedCustomer

type RemoveCustomerFunc func(customerId string)
