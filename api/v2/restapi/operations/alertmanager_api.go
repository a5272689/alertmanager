// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/prometheus/alertmanager/api/v2/restapi/operations/alert"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/alertgroups"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/general"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/receiver"
	"github.com/prometheus/alertmanager/api/v2/restapi/operations/silence"
)

// NewAlertmanagerAPI creates a new Alertmanager instance
func NewAlertmanagerAPI(spec *loads.Document) *AlertmanagerAPI {
	return &AlertmanagerAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		SilenceDeleteSilenceHandler: silence.DeleteSilenceHandlerFunc(func(params silence.DeleteSilenceParams) middleware.Responder {
			return middleware.NotImplemented("operation SilenceDeleteSilence has not yet been implemented")
		}),
		AlertgroupsGetAlertGroupsHandler: alertgroups.GetAlertGroupsHandlerFunc(func(params alertgroups.GetAlertGroupsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertgroupsGetAlertGroups has not yet been implemented")
		}),
		AlertGetAlertsHandler: alert.GetAlertsHandlerFunc(func(params alert.GetAlertsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertGetAlerts has not yet been implemented")
		}),
		ReceiverGetReceiversHandler: receiver.GetReceiversHandlerFunc(func(params receiver.GetReceiversParams) middleware.Responder {
			return middleware.NotImplemented("operation ReceiverGetReceivers has not yet been implemented")
		}),
		SilenceGetSilenceHandler: silence.GetSilenceHandlerFunc(func(params silence.GetSilenceParams) middleware.Responder {
			return middleware.NotImplemented("operation SilenceGetSilence has not yet been implemented")
		}),
		SilenceGetSilencesHandler: silence.GetSilencesHandlerFunc(func(params silence.GetSilencesParams) middleware.Responder {
			return middleware.NotImplemented("operation SilenceGetSilences has not yet been implemented")
		}),
		GeneralGetStatusHandler: general.GetStatusHandlerFunc(func(params general.GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation GeneralGetStatus has not yet been implemented")
		}),
		AlertPostAlertsHandler: alert.PostAlertsHandlerFunc(func(params alert.PostAlertsParams) middleware.Responder {
			return middleware.NotImplemented("operation AlertPostAlerts has not yet been implemented")
		}),
		SilencePostSilencesHandler: silence.PostSilencesHandlerFunc(func(params silence.PostSilencesParams) middleware.Responder {
			return middleware.NotImplemented("operation SilencePostSilences has not yet been implemented")
		}),
	}
}

/*AlertmanagerAPI API of the Prometheus Alertmanager (https://github.com/prometheus/alertmanager) */
type AlertmanagerAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// SilenceDeleteSilenceHandler sets the operation handler for the delete silence operation
	SilenceDeleteSilenceHandler silence.DeleteSilenceHandler
	// AlertgroupsGetAlertGroupsHandler sets the operation handler for the get alert groups operation
	AlertgroupsGetAlertGroupsHandler alertgroups.GetAlertGroupsHandler
	// AlertGetAlertsHandler sets the operation handler for the get alerts operation
	AlertGetAlertsHandler alert.GetAlertsHandler
	// ReceiverGetReceiversHandler sets the operation handler for the get receivers operation
	ReceiverGetReceiversHandler receiver.GetReceiversHandler
	// SilenceGetSilenceHandler sets the operation handler for the get silence operation
	SilenceGetSilenceHandler silence.GetSilenceHandler
	// SilenceGetSilencesHandler sets the operation handler for the get silences operation
	SilenceGetSilencesHandler silence.GetSilencesHandler
	// GeneralGetStatusHandler sets the operation handler for the get status operation
	GeneralGetStatusHandler general.GetStatusHandler
	// AlertPostAlertsHandler sets the operation handler for the post alerts operation
	AlertPostAlertsHandler alert.PostAlertsHandler
	// SilencePostSilencesHandler sets the operation handler for the post silences operation
	SilencePostSilencesHandler silence.PostSilencesHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AlertmanagerAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AlertmanagerAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AlertmanagerAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AlertmanagerAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AlertmanagerAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AlertmanagerAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AlertmanagerAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AlertmanagerAPI
func (o *AlertmanagerAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.SilenceDeleteSilenceHandler == nil {
		unregistered = append(unregistered, "silence.DeleteSilenceHandler")
	}

	if o.AlertgroupsGetAlertGroupsHandler == nil {
		unregistered = append(unregistered, "alertgroups.GetAlertGroupsHandler")
	}

	if o.AlertGetAlertsHandler == nil {
		unregistered = append(unregistered, "alert.GetAlertsHandler")
	}

	if o.ReceiverGetReceiversHandler == nil {
		unregistered = append(unregistered, "receiver.GetReceiversHandler")
	}

	if o.SilenceGetSilenceHandler == nil {
		unregistered = append(unregistered, "silence.GetSilenceHandler")
	}

	if o.SilenceGetSilencesHandler == nil {
		unregistered = append(unregistered, "silence.GetSilencesHandler")
	}

	if o.GeneralGetStatusHandler == nil {
		unregistered = append(unregistered, "general.GetStatusHandler")
	}

	if o.AlertPostAlertsHandler == nil {
		unregistered = append(unregistered, "alert.PostAlertsHandler")
	}

	if o.SilencePostSilencesHandler == nil {
		unregistered = append(unregistered, "silence.PostSilencesHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AlertmanagerAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AlertmanagerAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AlertmanagerAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AlertmanagerAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AlertmanagerAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AlertmanagerAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the alertmanager API
func (o *AlertmanagerAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AlertmanagerAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/silence/{silenceID}"] = silence.NewDeleteSilence(o.context, o.SilenceDeleteSilenceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alerts/groups"] = alertgroups.NewGetAlertGroups(o.context, o.AlertgroupsGetAlertGroupsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/alerts"] = alert.NewGetAlerts(o.context, o.AlertGetAlertsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/receivers"] = receiver.NewGetReceivers(o.context, o.ReceiverGetReceiversHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/silence/{silenceID}"] = silence.NewGetSilence(o.context, o.SilenceGetSilenceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/silences"] = silence.NewGetSilences(o.context, o.SilenceGetSilencesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/status"] = general.NewGetStatus(o.context, o.GeneralGetStatusHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/alerts"] = alert.NewPostAlerts(o.context, o.AlertPostAlertsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/silences"] = silence.NewPostSilences(o.context, o.SilencePostSilencesHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AlertmanagerAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AlertmanagerAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AlertmanagerAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AlertmanagerAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
