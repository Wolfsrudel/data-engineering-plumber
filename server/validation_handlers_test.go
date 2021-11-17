package server

import (
	"context"
	"io/ioutil"
	"sync"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	uuid "github.com/satori/go.uuid"
	"github.com/sirupsen/logrus"

	"github.com/batchcorp/plumber-schemas/build/go/protos"
	"github.com/batchcorp/plumber-schemas/build/go/protos/common"

	"github.com/batchcorp/plumber/config"
	"github.com/batchcorp/plumber/embed/etcd/etcdfakes"
	"github.com/batchcorp/plumber/validate"
)

var _ = Describe("Validation handlers", func() {

	var p *Server

	BeforeEach(func() {
		fakeEtcd := &etcdfakes.FakeIEtcd{}

		p = &Server{
			Etcd:      fakeEtcd,
			AuthToken: "batchcorp",
			PersistentConfig: &config.Config{
				ValidationsMutex: &sync.RWMutex{},
				SchemasMutex:     &sync.RWMutex{},
				Validations:      make(map[string]*common.Validation),
				Schemas:          make(map[string]*protos.Schema),
			},
			Log: logrus.NewEntry(&logrus.Logger{Out: ioutil.Discard}),
		}
	})

	Context("GetValidation", func() {
		It("checks auth token", func() {
			_, err := p.GetValidation(context.Background(), &protos.GetValidationRequest{
				Auth: &common.Auth{Token: "incorrect token"},
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrInvalidToken.Error()))
		})

		It("returns 404 when validation not found", func() {
			_, err := p.GetValidation(context.Background(), &protos.GetValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Id:   uuid.NewV4().String(),
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("validation not found"))
		})

		It("returns the validation", func() {
			v := &common.Validation{
				XId: uuid.NewV4().String(),
			}
			p.PersistentConfig.SetValidation(v.XId, v)

			resp, err := p.GetValidation(context.Background(), &protos.GetValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Id:   v.XId,
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Validation.XId).To(Equal(v.XId))
		})
	})

	Context("GetAllValidations", func() {
		It("checks auth token", func() {
			_, err := p.GetValidation(context.Background(), &protos.GetValidationRequest{
				Auth: &common.Auth{Token: "incorrect token"},
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrInvalidToken.Error()))
		})

		It("returns all validations", func() {
			v := &common.Validation{
				XId: uuid.NewV4().String(),
			}
			p.PersistentConfig.SetValidation(v.XId, v)

			resp, err := p.GetAllValidations(context.Background(), &protos.GetAllValidationsRequest{
				Auth: &common.Auth{Token: "batchcorp"},
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Validations[0].XId).To(Equal(v.XId))
		})
	})

	Context("CreateValidation", func() {
		It("checks auth token", func() {
			_, err := p.CreateValidation(context.Background(), &protos.CreateValidationRequest{
				Auth: &common.Auth{Token: "incorrect token"},
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrInvalidToken.Error()))
		})

		It("creates the validation and publishes CreateValidation message", func() {
			fakeEtcd := &etcdfakes.FakeIEtcd{}
			p.Etcd = fakeEtcd

			// ID is autogenerated, so we will check against schema ID
			schemaID := uuid.NewV4().String()

			resp, err := p.CreateValidation(context.Background(), &protos.CreateValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Validation: &common.Validation{
					SchemaId: schemaID,
				},
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Validation.SchemaId).To(Equal(schemaID))
			Expect(fakeEtcd.PutCallCount()).To(Equal(1))
			Expect(fakeEtcd.PublishCreateValidationCallCount()).To(Equal(1))
		})
	})

	Context("UpdateValidation", func() {
		It("checks auth token", func() {
			_, err := p.UpdateValidation(context.Background(), &protos.UpdateValidationRequest{
				Auth: &common.Auth{Token: "incorrect token"},
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrInvalidToken.Error()))
		})

		It("returns 404 when validation not found", func() {
			_, err := p.UpdateValidation(context.Background(), &protos.UpdateValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Id:   uuid.NewV4().String(),
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("validation not found"))
		})

		It("updates the validation", func() {
			fakeEtcd := &etcdfakes.FakeIEtcd{}
			p.Etcd = fakeEtcd

			v := &common.Validation{
				XId:      uuid.NewV4().String(),
				SchemaId: uuid.NewV4().String(),
			}

			p.PersistentConfig.SetValidation(v.XId, v)

			resp, err := p.UpdateValidation(context.Background(), &protos.UpdateValidationRequest{
				Auth:       &common.Auth{Token: "batchcorp"},
				Id:         v.XId,
				Validation: v,
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Validation.XId).To(Equal(v.XId))
			Expect(fakeEtcd.PublishUpdateValidationCallCount()).To(Equal(1))
			Expect(fakeEtcd.PutCallCount()).To(Equal(1))
		})
	})

	Context("DeleteValidation", func() {
		It("checks auth token", func() {
			_, err := p.DeleteValidation(context.Background(), &protos.DeleteValidationRequest{
				Auth: &common.Auth{Token: "incorrect token"},
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring(validate.ErrInvalidToken.Error()))
		})

		It("returns 404 when validation not found", func() {
			_, err := p.DeleteValidation(context.Background(), &protos.DeleteValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Id:   uuid.NewV4().String(),
			})

			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("validation not found"))
		})

		It("deletes a validation", func() {
			fakeEtcd := &etcdfakes.FakeIEtcd{}
			p.Etcd = fakeEtcd

			v := &common.Validation{
				XId:      uuid.NewV4().String(),
				SchemaId: uuid.NewV4().String(),
			}

			p.PersistentConfig.SetValidation(v.XId, v)

			resp, err := p.DeleteValidation(context.Background(), &protos.DeleteValidationRequest{
				Auth: &common.Auth{Token: "batchcorp"},
				Id:   v.XId,
			})

			Expect(err).ToNot(HaveOccurred())
			Expect(resp.Status.Code).To(Equal(common.Code_OK))
			Expect(len(p.PersistentConfig.Validations)).To(Equal(0))
			Expect(fakeEtcd.DeleteCallCount()).To(Equal(1))
			Expect(fakeEtcd.PublishDeleteValidationCallCount()).To(Equal(1))
		})
	})
})