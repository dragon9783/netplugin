// contivModelClient.go
// This file is auto generated by modelgen tool
// Do not edit this file manually

package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/contiv/objdb/modeldb"

	log "github.com/Sirupsen/logrus"
)

func httpGet(url string, jdata interface{}) error {

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode != int(200):
		log.Debugf("GET Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(response, jdata); err != nil {
		return err
	}

	return nil
}

func httpDelete(url string) error {

	req, err := http.NewRequest("DELETE", url, nil)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	// body, _ := ioutil.ReadAll(r.Body)

	switch {
	case r.StatusCode == int(404):
		// return errors.New("Page not found!")
		return nil
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode != int(200):
		log.Debugf("DELETE Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	return nil
}

func httpPost(url string, jdata interface{}) error {
	buf, err := json.Marshal(jdata)
	if err != nil {
		return err
	}

	body := bytes.NewBuffer(buf)
	r, err := http.Post(url, "application/json", body)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	switch {
	case r.StatusCode == int(404):
		return errors.New("Page not found!")
	case r.StatusCode == int(403):
		return errors.New("Access denied!")
	case r.StatusCode != int(200):
		log.Debugf("POST Status '%s' status code %d \n", r.Status, r.StatusCode)
		return errors.New(r.Status)
	}

	response, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	log.Debugf(string(response))

	return nil
}

// ContivClient has the contiv model client instance
type ContivClient struct {
	baseURL string
}

// NewContivClient returns a new client instance
func NewContivClient(baseURL string) (*ContivClient, error) {
	client := ContivClient{
		baseURL: baseURL,
	}

	return &client, nil
}

type App struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppName    string `json:"appName,omitempty"`    //
	TenantName string `json:"tenantName,omitempty"` //

	// add link-sets and links
	LinkSets AppLinkSets `json:"link-sets,omitempty"`
	Links    AppLinks    `json:"links,omitempty"`
}

type AppLinkSets struct {
	Services map[string]modeldb.Link `json:"Services,omitempty"`
}

type AppLinks struct {
	Tenant modeldb.Link `json:"Tenant,omitempty"`
}

type EndpointGroup struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	EndpointGroupID int      `json:"endpointGroupId,omitempty"` // Group Identifier
	GroupName       string   `json:"groupName,omitempty"`       // Group name
	NetworkName     string   `json:"networkName,omitempty"`     // Network
	Policies        []string `json:"policies,omitempty"`
	TenantName      string   `json:"tenantName,omitempty"` // Tenant

	// add link-sets and links
	LinkSets EndpointGroupLinkSets `json:"link-sets,omitempty"`
	Links    EndpointGroupLinks    `json:"links,omitempty"`
}

type EndpointGroupLinkSets struct {
	Policies map[string]modeldb.Link `json:"Policies,omitempty"`
	Services map[string]modeldb.Link `json:"Services,omitempty"`
}

type EndpointGroupLinks struct {
	Network modeldb.Link `json:"Network,omitempty"`
	Tenant  modeldb.Link `json:"Tenant,omitempty"`
}

type Global struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Name             string `json:"name,omitempty"`               // name of this block
	NetworkInfraType string `json:"network-infra-type,omitempty"` // Network infrastructure type
	Vlans            string `json:"vlans,omitempty"`              //
	Vxlans           string `json:"vxlans,omitempty"`             //

}

type Network struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Encap       string `json:"encap,omitempty"`       // Encapsulation
	Gateway     string `json:"gateway,omitempty"`     // Gateway
	IsPrivate   bool   `json:"isPrivate,omitempty"`   // Private network
	IsPublic    bool   `json:"isPublic,omitempty"`    // Public network
	NetworkName string `json:"networkName,omitempty"` // Network name
	PktTag      int    `json:"pktTag,omitempty"`      // Vlan/Vxlan Tag
	Subnet      string `json:"subnet,omitempty"`      // Subnet
	TenantName  string `json:"tenantName,omitempty"`  // Tenant Name

	// add link-sets and links
	LinkSets NetworkLinkSets `json:"link-sets,omitempty"`
	Links    NetworkLinks    `json:"links,omitempty"`
}

type NetworkLinkSets struct {
	EndpointGroups map[string]modeldb.Link `json:"EndpointGroups,omitempty"`
	Services       map[string]modeldb.Link `json:"Services,omitempty"`
}

type NetworkLinks struct {
	Tenant modeldb.Link `json:"Tenant,omitempty"`
}

type Policy struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	PolicyName string `json:"policyName,omitempty"` // Policy Name
	TenantName string `json:"tenantName,omitempty"` // Tenant Name

	// add link-sets and links
	LinkSets PolicyLinkSets `json:"link-sets,omitempty"`
	Links    PolicyLinks    `json:"links,omitempty"`
}

type PolicyLinkSets struct {
	EndpointGroups map[string]modeldb.Link `json:"EndpointGroups,omitempty"`
	Rules          map[string]modeldb.Link `json:"Rules,omitempty"`
}

type PolicyLinks struct {
	Tenant modeldb.Link `json:"Tenant,omitempty"`
}

type Rule struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	Action        string `json:"action,omitempty"`        // Action
	Direction     string `json:"direction,omitempty"`     // Direction
	EndpointGroup string `json:"endpointGroup,omitempty"` // Group
	IpAddress     string `json:"ipAddress,omitempty"`     // IP Address
	Network       string `json:"network,omitempty"`       // Network Name
	PolicyName    string `json:"policyName,omitempty"`    // Policy Name
	Port          int    `json:"port,omitempty"`          // Port No
	Priority      int    `json:"priority,omitempty"`      // Priority
	Protocol      string `json:"protocol,omitempty"`      // Protocol
	RuleID        string `json:"ruleId,omitempty"`        // Rule Id
	TenantName    string `json:"tenantName,omitempty"`    // Tenant Name

	// add link-sets and links
	LinkSets RuleLinkSets `json:"link-sets,omitempty"`
}

type RuleLinkSets struct {
	Policies map[string]modeldb.Link `json:"Policies,omitempty"`
}

type Service struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppName        string   `json:"appName,omitempty"` //
	Command        string   `json:"command,omitempty"` //
	Cpu            string   `json:"cpu,omitempty"`     //
	EndpointGroups []string `json:"endpointGroups,omitempty"`
	Environment    []string `json:"environment,omitempty"`
	ImageName      string   `json:"imageName,omitempty"` //
	Memory         string   `json:"memory,omitempty"`    //
	Networks       []string `json:"networks,omitempty"`
	Scale          int      `json:"scale,omitempty"`         //
	ServiceName    string   `json:"serviceName,omitempty"`   //
	TenantName     string   `json:"tenantName,omitempty"`    //
	VolumeProfile  string   `json:"volumeProfile,omitempty"` //

	// add link-sets and links
	LinkSets ServiceLinkSets `json:"link-sets,omitempty"`
	Links    ServiceLinks    `json:"links,omitempty"`
}

type ServiceLinkSets struct {
	EndpointGroups map[string]modeldb.Link `json:"EndpointGroups,omitempty"`
	Instances      map[string]modeldb.Link `json:"Instances,omitempty"`
	Networks       map[string]modeldb.Link `json:"Networks,omitempty"`
}

type ServiceLinks struct {
	App           modeldb.Link `json:"App,omitempty"`
	VolumeProfile modeldb.Link `json:"VolumeProfile,omitempty"`
}

type ServiceInstance struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	AppName     string   `json:"appName,omitempty"`     //
	InstanceID  string   `json:"instanceId,omitempty"`  //
	ServiceName string   `json:"serviceName,omitempty"` //
	TenantName  string   `json:"tenantName,omitempty"`  //
	Volumes     []string `json:"volumes,omitempty"`

	// add link-sets and links
	LinkSets ServiceInstanceLinkSets `json:"link-sets,omitempty"`
	Links    ServiceInstanceLinks    `json:"links,omitempty"`
}

type ServiceInstanceLinkSets struct {
	Volumes map[string]modeldb.Link `json:"Volumes,omitempty"`
}

type ServiceInstanceLinks struct {
	Service modeldb.Link `json:"Service,omitempty"`
}

type Tenant struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DefaultNetwork string `json:"defaultNetwork,omitempty"` // Network name
	TenantName     string `json:"tenantName,omitempty"`     // Tenant Name

	// add link-sets and links
	LinkSets TenantLinkSets `json:"link-sets,omitempty"`
}

type TenantLinkSets struct {
	Apps           map[string]modeldb.Link `json:"Apps,omitempty"`
	EndpointGroups map[string]modeldb.Link `json:"EndpointGroups,omitempty"`
	Networks       map[string]modeldb.Link `json:"Networks,omitempty"`
	Policies       map[string]modeldb.Link `json:"Policies,omitempty"`
	VolumeProfiles map[string]modeldb.Link `json:"VolumeProfiles,omitempty"`
	Volumes        map[string]modeldb.Link `json:"Volumes,omitempty"`
}

type Volume struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType string `json:"datastoreType,omitempty"` //
	MountPoint    string `json:"mountPoint,omitempty"`    //
	PoolName      string `json:"poolName,omitempty"`      //
	Size          string `json:"size,omitempty"`          //
	TenantName    string `json:"tenantName,omitempty"`    //
	VolumeName    string `json:"volumeName,omitempty"`    //

	// add link-sets and links
	LinkSets VolumeLinkSets `json:"link-sets,omitempty"`
	Links    VolumeLinks    `json:"links,omitempty"`
}

type VolumeLinkSets struct {
	ServiceInstances map[string]modeldb.Link `json:"ServiceInstances,omitempty"`
}

type VolumeLinks struct {
	Tenant modeldb.Link `json:"Tenant,omitempty"`
}

type VolumeProfile struct {
	// every object has a key
	Key string `json:"key,omitempty"`

	DatastoreType     string `json:"datastoreType,omitempty"`     //
	MountPoint        string `json:"mountPoint,omitempty"`        //
	PoolName          string `json:"poolName,omitempty"`          //
	Size              string `json:"size,omitempty"`              //
	TenantName        string `json:"tenantName,omitempty"`        //
	VolumeProfileName string `json:"volumeProfileName,omitempty"` //

	// add link-sets and links
	LinkSets VolumeProfileLinkSets `json:"link-sets,omitempty"`
	Links    VolumeProfileLinks    `json:"links,omitempty"`
}

type VolumeProfileLinkSets struct {
	Services map[string]modeldb.Link `json:"Services,omitempty"`
}

type VolumeProfileLinks struct {
	Tenant modeldb.Link `json:"Tenant,omitempty"`
}

// AppPost posts the app object
func (c *ContivClient) AppPost(obj *App) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppName
	url := c.baseURL + "/api/apps/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating app %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// AppGet gets the app object
func (c *ContivClient) AppGet(tenantName string, appName string) (*App, error) {
	// build key and URL
	keyStr := tenantName + ":" + appName
	url := c.baseURL + "/api/apps/" + keyStr + "/"

	// http get the object
	var obj App
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting app %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// AppDelete deletes the app object
func (c *ContivClient) AppDelete(tenantName string, appName string) error {
	// build key and URL
	keyStr := tenantName + ":" + appName
	url := c.baseURL + "/api/apps/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting app %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// EndpointGroupPost posts the endpointGroup object
func (c *ContivClient) EndpointGroupPost(obj *EndpointGroup) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName + ":" + obj.GroupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating endpointGroup %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// EndpointGroupGet gets the endpointGroup object
func (c *ContivClient) EndpointGroupGet(tenantName string, networkName string, groupName string) (*EndpointGroup, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + groupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http get the object
	var obj EndpointGroup
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting endpointGroup %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// EndpointGroupDelete deletes the endpointGroup object
func (c *ContivClient) EndpointGroupDelete(tenantName string, networkName string, groupName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName + ":" + groupName
	url := c.baseURL + "/api/endpointGroups/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting endpointGroup %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// GlobalPost posts the global object
func (c *ContivClient) GlobalPost(obj *Global) error {
	// build key and URL
	keyStr := obj.Name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating global %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// GlobalGet gets the global object
func (c *ContivClient) GlobalGet(name string) (*Global, error) {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http get the object
	var obj Global
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting global %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// GlobalDelete deletes the global object
func (c *ContivClient) GlobalDelete(name string) error {
	// build key and URL
	keyStr := name
	url := c.baseURL + "/api/globals/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting global %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// NetworkPost posts the network object
func (c *ContivClient) NetworkPost(obj *Network) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.NetworkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating network %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// NetworkGet gets the network object
func (c *ContivClient) NetworkGet(tenantName string, networkName string) (*Network, error) {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http get the object
	var obj Network
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting network %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// NetworkDelete deletes the network object
func (c *ContivClient) NetworkDelete(tenantName string, networkName string) error {
	// build key and URL
	keyStr := tenantName + ":" + networkName
	url := c.baseURL + "/api/networks/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting network %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// PolicyPost posts the policy object
func (c *ContivClient) PolicyPost(obj *Policy) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating policy %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// PolicyGet gets the policy object
func (c *ContivClient) PolicyGet(tenantName string, policyName string) (*Policy, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http get the object
	var obj Policy
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting policy %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// PolicyDelete deletes the policy object
func (c *ContivClient) PolicyDelete(tenantName string, policyName string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName
	url := c.baseURL + "/api/policys/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting policy %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// RulePost posts the rule object
func (c *ContivClient) RulePost(obj *Rule) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.PolicyName + ":" + obj.RuleID
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating rule %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// RuleGet gets the rule object
func (c *ContivClient) RuleGet(tenantName string, policyName string, ruleId string) (*Rule, error) {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http get the object
	var obj Rule
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting rule %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// RuleDelete deletes the rule object
func (c *ContivClient) RuleDelete(tenantName string, policyName string, ruleId string) error {
	// build key and URL
	keyStr := tenantName + ":" + policyName + ":" + ruleId
	url := c.baseURL + "/api/rules/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting rule %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ServicePost posts the service object
func (c *ContivClient) ServicePost(obj *Service) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppName + ":" + obj.ServiceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating service %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ServiceGet gets the service object
func (c *ContivClient) ServiceGet(tenantName string, appName string, serviceName string) (*Service, error) {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http get the object
	var obj Service
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting service %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceDelete deletes the service object
func (c *ContivClient) ServiceDelete(tenantName string, appName string, serviceName string) error {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName
	url := c.baseURL + "/api/services/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting service %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// ServiceInstancePost posts the serviceInstance object
func (c *ContivClient) ServiceInstancePost(obj *ServiceInstance) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.AppName + ":" + obj.ServiceName + ":" + obj.InstanceID
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating serviceInstance %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// ServiceInstanceGet gets the serviceInstance object
func (c *ContivClient) ServiceInstanceGet(tenantName string, appName string, serviceName string, instanceId string) (*ServiceInstance, error) {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName + ":" + instanceId
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http get the object
	var obj ServiceInstance
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting serviceInstance %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// ServiceInstanceDelete deletes the serviceInstance object
func (c *ContivClient) ServiceInstanceDelete(tenantName string, appName string, serviceName string, instanceId string) error {
	// build key and URL
	keyStr := tenantName + ":" + appName + ":" + serviceName + ":" + instanceId
	url := c.baseURL + "/api/serviceInstances/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting serviceInstance %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// TenantPost posts the tenant object
func (c *ContivClient) TenantPost(obj *Tenant) error {
	// build key and URL
	keyStr := obj.TenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating tenant %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// TenantGet gets the tenant object
func (c *ContivClient) TenantGet(tenantName string) (*Tenant, error) {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http get the object
	var obj Tenant
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting tenant %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// TenantDelete deletes the tenant object
func (c *ContivClient) TenantDelete(tenantName string) error {
	// build key and URL
	keyStr := tenantName
	url := c.baseURL + "/api/tenants/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting tenant %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumePost posts the volume object
func (c *ContivClient) VolumePost(obj *Volume) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating volume %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeGet gets the volume object
func (c *ContivClient) VolumeGet(tenantName string, volumeName string) (*Volume, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http get the object
	var obj Volume
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting volume %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeDelete deletes the volume object
func (c *ContivClient) VolumeDelete(tenantName string, volumeName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeName
	url := c.baseURL + "/api/volumes/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting volume %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}

// VolumeProfilePost posts the volumeProfile object
func (c *ContivClient) VolumeProfilePost(obj *VolumeProfile) error {
	// build key and URL
	keyStr := obj.TenantName + ":" + obj.VolumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http post the object
	err := httpPost(url, obj)
	if err != nil {
		log.Errorf("Error creating volumeProfile %+v. Err: %v", obj, err)
		return err
	}

	return nil
}

// VolumeProfileGet gets the volumeProfile object
func (c *ContivClient) VolumeProfileGet(tenantName string, volumeProfileName string) (*VolumeProfile, error) {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http get the object
	var obj VolumeProfile
	err := httpGet(url, &obj)
	if err != nil {
		log.Errorf("Error getting volumeProfile %+v. Err: %v", keyStr, err)
		return nil, err
	}

	return &obj, nil
}

// VolumeProfileDelete deletes the volumeProfile object
func (c *ContivClient) VolumeProfileDelete(tenantName string, volumeProfileName string) error {
	// build key and URL
	keyStr := tenantName + ":" + volumeProfileName
	url := c.baseURL + "/api/volumeProfiles/" + keyStr + "/"

	// http get the object
	err := httpDelete(url)
	if err != nil {
		log.Errorf("Error deleting volumeProfile %s. Err: %v", keyStr, err)
		return err
	}

	return nil
}
