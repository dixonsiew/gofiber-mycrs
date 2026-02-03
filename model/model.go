package model

import "encoding/xml"

type Setting struct {
    DB_URL          string
    DB_PORT         string
    DB_SERVER       string
    DB_SERVICE      string
    DB_USER         string
    DB_PASSWORD     string
    PORT            string
    BASE_URL        string
    MAIL_HOST       string
    MAIL_PORT       string
    MAIL_USERNAME   string
    MAIL_PASSWORD   string
    MAIL_REQUIRETLS string
    MAIL_SENDER     string
    MAIL_FROM_NAME  string
    MAIL_APP_NAME   string
    MAIL_TO_ADDRESS string
    PUSH_NTFY       string
}

type Fault struct {
    XMLName     xml.Name `xml:"Fault" json:"-"`
    Faultcode   string   `xml:"faultcode" json:"faultcode"`
    Faultstring string   `xml:"faultstring" json:"faultstring"`
}

type EnrollRoamingCertResponse struct {
    XMLName                 xml.Name                `xml:"EnrollRoamingCertResponse"`
    EnrollRoamingCertResult EnrollRoamingCertResult `xml:"EnrollRoamingCertResult"`
}

type EnrollRoamingCertResult struct {
    XMLName       xml.Name `xml:"EnrollRoamingCertResult" json:"-"`
    CertInBased64 string   `xml:"CertInBased64" json:"CertInBased64"`
    StatusCode    string   `xml:"StatusCode" json:"StatusCode"`
    StatusMessage string   `xml:"StatusMessage" json:"StatusMessage"`
    CertSerialNo  string   `xml:"CertSerialNo" json:"CertSerialNo"`
    RequestCode   string   `xml:"RequestCode" json:"RequestCode"`
    ValidFrom     string   `xml:"ValidFrom" json:"ValidFrom"`
    ValidTo       string   `xml:"ValidTo"`
}

type GenerateResetPinCodeResponse struct {
    XMLName                    xml.Name                   `xml:"GenerateResetPinCodeResponse"`
    GenerateResetPinCodeResult GenerateResetPinCodeResult `xml:"GenerateResetPinCodeResult"`
}

type GenerateResetPinCodeResult struct {
    XMLName       xml.Name `xml:"GenerateResetPinCodeResult" json:"-"`
    ResetCode     string   `xml:"ResetCode" json:"ResetCode"`
    StatusCode    string   `xml:"StatusCode" json:"StatusCode"`
    StatusMessage string   `xml:"StatusMessage" json:"StatusMessage"`
}

type ResetPinResponse struct {
    XMLName        xml.Name       `xml:"ResetPinResponse"`
    ResetPinResult ResetPinResult `xml:"ResetPinResult"`
}

type ResetPinResult struct {
    XMLName       xml.Name `xml:"ResetPinResult" json:"-"`
    ResetCode     string   `xml:"ResetCode" json:"ResetCode"`
    StatusCode    string   `xml:"StatusCode" json:"StatusCode"`
    StatusMessage string   `xml:"StatusMessage" json:"StatusMessage"`
}

type RequestRevokeResponse struct {
    XMLName             xml.Name            `xml:"RequestRevokeResponse"`
    RequestRevokeResult RequestRevokeResult `xml:"RequestRevokeResult"`
}

type RequestRevokeResult struct {
    XMLName       xml.Name `xml:"RequestRevokeResult" json:"-"`
    ResetCode     string   `xml:"ResetCode" json:"ResetCode"`
    StatusCode    string   `xml:"StatusCode" json:"StatusCode"`
    StatusMessage string   `xml:"StatusMessage" json:"StatusMessage"`
}

type CertificateStatusIIHResponse struct {
    XMLName                    xml.Name                   `xml:"certificateStatusIIHResponse"`
    CertificateStatusIIHResult CertificateStatusIIHResult `xml:"certificateStatusIIHResult"`
}

type CertificateStatusIIHResult struct {
    XMLName     xml.Name `xml:"certificateStatusIIHResult" json:"-"`
    Cert        Cert     `xml:"Cert"`
    Recordcount string   `xml:"Recordcount" json:"Recordcount"`
    Statuscode  string   `xml:"Statuscode"`
}

type Cert struct {
    XMLName  xml.Name   `xml:"Cert" json:"-"`
    CertList []CertList `xml:"CertList"`
}

type CertList struct {
    XMLName     xml.Name `xml:"CertList" json:"-"`
    IcNo        string   `xml:"IcNo" json:"IcNo"`
    RequestCode string   `xml:"RequestCode" json:"RequestCode"`
    Status      string   `xml:"Status" json:"Status"`
    TraderId    string   `xml:"TraderId" json:"TraderId"`
    ValidFrom   string   `xml:"ValidFrom" json:"ValidFrom"`
    ValidTo     string   `xml:"ValidTo" json:"ValidTo"`
}
