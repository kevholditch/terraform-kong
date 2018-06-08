package gokong

import (
	"encoding/json"
	"fmt"
)

type CertificateClient struct {
	config *Config
}

type CertificateRequest struct {
	Cert *string `json:"cert,omitempty"`
	Key  *string `json:"key,omitempty"`
}

type Certificate struct {
	Id   *string `json:"id,omitempty"`
	Cert *string `json:"cert,omitempty"`
	Key  *string `json:"key,omitempty"`
}

type Certificates struct {
	Results []*Certificate `json:"data,omitempty"`
	Total   int            `json:"total,omitempty"`
}

const CertificatesPath = "/certificates/"

func (certificateClient *CertificateClient) GetById(id string) (*Certificate, error) {

	_, body, errs := NewRequest(certificateClient.config).Get(certificateClient.config.HostAddress + CertificatesPath + id).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get certificate, error: %v", errs)
	}

	certificate := &Certificate{}
	err := json.Unmarshal([]byte(body), certificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate get response, error: %v", err)
	}

	if certificate.Id == nil {
		return nil, nil
	}

	return certificate, nil
}

func (certificateClient *CertificateClient) Create(certificateRequest *CertificateRequest) (*Certificate, error) {

	_, body, errs := NewRequest(certificateClient.config).Post(certificateClient.config.HostAddress + CertificatesPath).Send(certificateRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not create new certificate, error: %v", errs)
	}

	createdCertificate := &Certificate{}
	err := json.Unmarshal([]byte(body), createdCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate creation response, error: %v", err)
	}

	if createdCertificate.Id == nil {
		return nil, fmt.Errorf("could not create certificate, error: %v", body)
	}

	return createdCertificate, nil
}

func (certificateClient *CertificateClient) DeleteById(id string) error {

	res, _, errs := NewRequest(certificateClient.config).Delete(certificateClient.config.HostAddress + CertificatesPath + id).End()
	if errs != nil {
		return fmt.Errorf("could not delete certificate, result: %v error: %v", res, errs)
	}

	return nil
}

func (certificateClient *CertificateClient) List() (*Certificates, error) {

	_, body, errs := NewRequest(certificateClient.config).Get(certificateClient.config.HostAddress + CertificatesPath).End()
	if errs != nil {
		return nil, fmt.Errorf("could not get certificates, error: %v", errs)
	}

	certificates := &Certificates{}
	err := json.Unmarshal([]byte(body), certificates)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificates list response, error: %v", err)
	}

	return certificates, nil
}

func (certificateClient *CertificateClient) UpdateById(id string, certificateRequest *CertificateRequest) (*Certificate, error) {

	_, body, errs := NewRequest(certificateClient.config).Patch(certificateClient.config.HostAddress + CertificatesPath + id).Send(certificateRequest).End()
	if errs != nil {
		return nil, fmt.Errorf("could not update certificate, error: %v", errs)
	}

	updatedCertificate := &Certificate{}
	err := json.Unmarshal([]byte(body), updatedCertificate)
	if err != nil {
		return nil, fmt.Errorf("could not parse certificate update response, error: %v", err)
	}

	if updatedCertificate.Id == nil {
		return nil, fmt.Errorf("could not update certificate, error: %v", body)
	}

	return updatedCertificate, nil
}
