package utils_test

import (
	"context"

	"code.cloudfoundry.org/brokerapi/v13/domain"
	"code.cloudfoundry.org/brokerapi/v13/utils"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Context", func() {
	type testContextKey string

	var (
		ctx                   context.Context
		contextValidatorKey   testContextKey
		contextValidatorValue string
	)

	BeforeEach(func() {
		contextValidatorKey = "context-utilities-test"
		contextValidatorValue = "original"
		ctx = context.Background()
		ctx = context.WithValue(ctx, contextValidatorKey, contextValidatorValue)
	})

	Describe("Service Context", func() {
		Context("when the service is nil", func() {
			It("returns the original context", func() {
				ctx = utils.AddServiceToContext(ctx, nil)
				Expect(ctx.Err()).To(BeZero())
				Expect(utils.RetrieveServiceFromContext(ctx)).To(BeZero())
				Expect(ctx.Value(contextValidatorKey).(string)).To(Equal(contextValidatorValue))
			})
		})

		Context("when the service is valid", func() {
			It("sets and receives the service in the context", func() {
				service := &domain.Service{
					ID:   "9A3095D7-ED3C-45FA-BC9F-592820628723",
					Name: "Test Service",
				}
				ctx = utils.AddServiceToContext(ctx, service)
				Expect(ctx.Err()).To(BeZero())
				Expect(ctx.Value(contextValidatorKey).(string)).To(Equal(contextValidatorValue))
				Expect(utils.RetrieveServiceFromContext(ctx).ID).To(Equal(service.ID))
				Expect(utils.RetrieveServiceFromContext(ctx).Name).To(Equal(service.Name))
				Expect(utils.RetrieveServiceFromContext(ctx).Metadata).To(BeZero())
			})
		})
	})

	Describe("Plan Context", func() {
		Context("when the service plan is nil", func() {
			It("returns the original context", func() {
				ctx = utils.AddServicePlanToContext(ctx, nil)
				Expect(ctx.Err()).To(BeZero())
				Expect(utils.RetrieveServicePlanFromContext(ctx)).To(BeZero())
				Expect(ctx.Value(contextValidatorKey).(string)).To(Equal(contextValidatorValue))
			})
		})

		Context("when the service plan is valid", func() {
			It("sets and retrieves the service plan in the context", func() {
				plan := &domain.ServicePlan{
					ID: "AC257573-8C62-4B1A-AC34-ECA3863F50EC",
				}
				ctx = utils.AddServicePlanToContext(ctx, plan)
				Expect(ctx.Err()).To(BeZero())
				Expect(ctx.Value(contextValidatorKey).(string)).To(Equal(contextValidatorValue))
				Expect(utils.RetrieveServicePlanFromContext(ctx).ID).To(Equal(plan.ID))
			})
		})
	})
})
