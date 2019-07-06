package storage

import (
	"customer-app/internal/adding"
	"customer-app/internal/retrieving"
	"customer-app/internal/updating"
	"customer-app/internal/uuid"
)

type InsertCustomerFunc func(unPersistedCustomer adding.UnPersistedCustomer, genUUIDStr uuid.GenFunc) adding.PersistedCustomer

type RetrieveCustomerFunc func(customerId string) retrieving.Customer

type RetrieveCustomersFunc func(request retrieving.MultiRequest) []retrieving.Customer

type UpdateCustomerFunc func(request updating.Request) updating.UpdatedCustomer

type RemoveCustomerFunc func(customerId string)
