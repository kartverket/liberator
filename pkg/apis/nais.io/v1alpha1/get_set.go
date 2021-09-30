package nais_io_v1alpha1

import nais_io_v1 "github.com/nais/liberator/pkg/apis/nais.io/v1"

// TODO: replace manual getters with generated code
// TODO: candidates are either `go generate` or a switch to Protobuf

func (in *Application) SetStatus(status *nais_io_v1.Status) {
	in.Status = *status
}

func (in *Application) GetStatus() *nais_io_v1.Status {
	return &in.Status
}

func (in *Application) GetStrategy() *nais_io_v1.Strategy {
	return in.Spec.Strategy
}

func (in *Application) GetReplicas() *nais_io_v1.Replicas {
	return in.Spec.Replicas
}

func (in *Application) GetCleanup() *nais_io_v1.Cleanup {
	return in.Spec.Cleanup
}

func (in *Application) GetPrometheus() *nais_io_v1.PrometheusConfig {
	return in.Spec.Prometheus
}

func (in *Application) GetLogtransform() string {
	return in.Spec.Logtransform
}

func (in *Application) GetLogformat() string {
	return in.Spec.Logformat
}

func (in *Application) GetPort() int {
	return in.Spec.Port
}

func (in *Application) GetService() *nais_io_v1.Service {
	return in.Spec.Service
}

func (in *Application) GetLiveness() *nais_io_v1.Probe {
	return in.Spec.Liveness
}

func (in *Application) GetReadiness() *nais_io_v1.Probe {
	return in.Spec.Readiness
}

func (in *Application) GetStartup() *nais_io_v1.Probe {
	return in.Spec.Startup
}

func (in *Application) GetEnv() nais_io_v1.EnvVars {
	return in.Spec.Env
}

func (in *Application) GetImage() string {
	return in.Spec.Image
}

func (in *Application) GetCommand() []string {
	return in.Spec.Command
}

func (in *Application) GetFilesFrom() []nais_io_v1.FilesFrom {
	return in.Spec.FilesFrom
}

func (in *Application) GetEnvFrom() []nais_io_v1.EnvFrom {
	return in.Spec.EnvFrom
}

func (in *Application) GetPreStopHook() *nais_io_v1.PreStopHook {
	return in.Spec.PreStopHook
}

func (in *Application) GetPreStopHookPath() string {
	return in.Spec.PreStopHookPath
}

func (in *Application) GetResources() *nais_io_v1.ResourceRequirements {
	return in.Spec.Resources
}