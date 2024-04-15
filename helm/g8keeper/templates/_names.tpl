{{/*
Tinksrv component names
*/}}

{{- define "g8keeper.tinksrv.svc.name" -}}
{{- printf "svc-tinksrv-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.tinksrv.dplm.name" -}}
{{- printf "dplm-tinksrv-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.tinksrv.cm.name" -}}
{{- printf "cm-tinksrv-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.tinksrv.sct.name" -}}
{{- printf "sct-tinksrv-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.tinksrv.sct.mountPath" -}}
/run/secrets
{{- end }}
{{- define "g8keeper.tinksrv.sct.fileName" -}}
keyset.json
{{- end }}

{{/*
Backend component names
*/}}

{{- define "g8keeper.backend.svc.name" -}}
{{- printf "svc-backend-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.backend.dplm.name" -}}
{{- printf "dplm-backend-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}

{{- define "g8keeper.backend.cm.name" -}}
{{- printf "cm-backend-%s" (include "g8keeper.fullname" .) | trunc 63 | trimSuffix "-" }}
{{- end }}
