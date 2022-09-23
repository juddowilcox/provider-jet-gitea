{{/*
Expand the name of the chart.
*/}}

{{- define "provider-jet-gitea.name" -}}
{{- default .Chart.Name .Values.nameOverride | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Create a default fully qualified app name.
We truncate at 63 chars because some Kubernetes name fields are limited to this (by the DNS naming spec).
If release name contains chart name it will be used as a full name.
*/}}

{{- define "provider-jet-gitea.fullname" -}}
{{- if .Values.fullnameOverride }}
{{- .Values.fullnameOverride | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- $name := default .Chart.Name .Values.nameOverride }}
{{- if contains $name .Release.Name }}
{{- .Release.Name | trunc 63 | trimSuffix "-" }}
{{- else }}
{{- printf "%s-%s" .Release.Name $name | trunc 63 | trimSuffix "-" }}
{{- end }}
{{- end }}
{{- end }}

{{/*
Create chart name and version as used by the chart label.
*/}}

{{- define "provider-jet-gitea.chart" -}}
{{- printf "%s-%s" .Chart.Name .Chart.Version | replace "+" "_" | trunc 63 | trimSuffix "-" }}
{{- end }}

{{/*
Globals
*/}}

{{- define "global.registry" -}}
{{- if .Values.global.registry -}}
{{ .Values.global.registry | trimSuffix "/" }}/
{{- end -}}
{{- end -}}

{{/*
Configs
*/}}

{{- define "controllerconfig.image.repository" -}}
{{- if .Values.controllerconfig.image.repository -}}
{{ .Values.controllerconfig.image.repository }}
{{- else -}}
crossplane
{{- end -}}
{{- end -}}

{{- define "controllerconfig.image.name" -}}
{{- if .Values.controllerconfig.image.name -}}
{{ .Values.controllerconfig.image.name }}
{{- else -}}
provider-jet-gitea-controller
{{- end -}}
{{- end -}}

{{- define "controllerconfig.image.tag" -}}
{{- if .Values.controllerconfig.image.tag -}}
{{ .Values.controllerconfig.image.tag }}
{{- else -}}
{{ .Chart.AppVersion }}
{{- end -}}
{{- end -}}

{{- define "controllerconfig.args" -}}
{{- if .Values.controllerconfig.args }}
args:
{{ toYaml ( .Values.controllerconfig.args ) }}
{{- end -}}
{{- end -}}

{{- define "image.pullPolicy" -}}
{{- if .Values.image.pullPolicy }}
imagePullPolicy: {{ .Values.image.pullPolicy -}}
{{- end -}}
{{- end -}}

{{- define "replicas" -}}
{{- if .Values.replicas }}
replicas: {{ .Values.replicas -}}
{{- end -}}
{{- end -}}

{{- define "serviceAccount.name" }}
serviceAccountName: {{ include "provider-jet-gitea.serviceAccountName" . -}}
{{- end -}}

{{- define "provider.image.repository" -}}
{{- if .Values.provider.image.repository -}}
{{ .Values.provider.image.repository }}
{{- else -}}
crossplane
{{- end -}}
{{- end -}}

{{- define "provider.image.name" -}}
{{- if .Values.provider.image.name -}}
{{ .Values.provider.image.name }}
{{- else -}}
provider-jet-gitea
{{- end -}}
{{- end -}}

{{- define "provider.image.tag" -}}
{{- if .Values.provider.image.tag -}}
{{ .Values.provider.image.tag }}
{{- else -}}
{{ .Chart.AppVersion }}
{{- end -}}
{{- end -}}

{{- define "affinity" -}}
{{- if .Values.affinity }}
affinity:
{{ toYaml ( .Values.affinity ) | indent 2 }}
{{- end -}}
{{- end -}}

{{- define "imagePullSecrets" -}}
{{- if .Values.image.pullSecrets }}
imagePullSecrets:
{{ toYaml ( .Values.image.pullSecrets ) }}
{{- end -}}
{{- end -}}

{{- define "nodeSelector" -}}
{{- if .Values.nodeSelector }}
nodeSelector:
{{ toYaml ( .Values.nodeSelector ) | indent 2 }}
{{- end -}}
{{- end -}}

{{- define "podSecurityContext" -}}
{{- if .Values.podSecurityContext }}
podSecurityContext:
{{ toYaml ( .Values.podSecurityContext ) | indent 2 }}
{{- end -}}
{{- end -}}

{{- define "resources" -}}
{{- if .Values.resources }}
resources:
{{ toYaml ( .Values.resources ) | indent 2 }}
{{- end -}}
{{- end -}}

{{- define "tolerations" -}}
{{- if .Values.tolerations }}
tolerations:
{{ toYaml ( .Values.tolerations ) }}
{{- end -}}
{{- end -}}

{{- define "securityContext" -}}
{{- if .Values.securityContext }}
securityContext:
{{ toYaml ( .Values.securityContext ) | indent 2 }}
{{- end -}}
{{- end -}}

{{/*
Common labels
*/}}
{{- define "provider-jet-gitea.labels" -}}
helm.sh/chart: {{ include "provider-jet-gitea.chart" . }}
{{ include "provider-jet-gitea.selectorLabels" . }}
{{- if .Chart.AppVersion }}
app.kubernetes.io/version: {{ .Chart.AppVersion | quote }}
{{- end }}
app.kubernetes.io/managed-by: {{ .Release.Service }}
{{- end }}

{{/*
Selector labels
*/}}
{{- define "provider-jet-gitea.selectorLabels" -}}
app.kubernetes.io/name: {{ include "provider-jet-gitea.name" . }}
app.kubernetes.io/instance: {{ .Release.Name }}
{{- end }}

{{/*
Create the name of the service account to use
*/}}
{{- define "provider-jet-gitea.serviceAccountName" -}}
{{- .Values.serviceAccount.name | default "crossplane" }}
{{- end }}
