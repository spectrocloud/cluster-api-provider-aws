---
skill: Upgrade Go Version - FIPS and Non-FIPS Build Pipeline
description: Procedure for upgrading the Go version used across the spectro release workflow and Makefile. Covers BUILDER_GOLANG_VERSION and GO_VERSION for both FIPS and non-FIPS Docker image builds, the Docker build-arg pipeline from workflow to base image, and Makefile defaults. A single target version is applied to all 6 locations to keep the build pipeline consistent.
type: procedure
repository: cluster-api-provider-aws
team: cloud
topics: [go, upgrade, docker, fips, ci, github-actions, makefile, build]
difficulty: low
last_updated: 2026-02-25
related_skills: []
memory_references: []
---


# Upgrade Go Version - FIPS and Non-FIPS Build Pipeline

## Overview

Procedure for upgrading the Go version across the CAPA (Cluster API Provider AWS) build pipeline. The Go version is referenced in 6 locations across 2 files and must be kept in sync for both FIPS and non-FIPS Docker image builds.

**Complexity**: Low
**Prerequisites**: Knowledge of which Go version to target

> **Note**: The same Go version is used for both FIPS and non-FIPS builds. There is no scenario where these should differ.

## Why This Procedure Exists

The CAPA build pipeline uses Go in two layers:

1. **Docker Build Image** (`BUILDER_GOLANG_VERSION`): Controls the base image used to compile the controller binary inside Docker. This is passed as a Docker build-arg and resolves to a Spectro Cloud managed base image.
2. **CI Go Version** (`GO_VERSION`): Controls the Go version used in the CI container for non-Docker build tasks.

Both variables are set in the GitHub Actions workflow for release builds and have fallback defaults in the Makefile for local development. When Go is upgraded, all 6 references must be updated atomically to prevent version drift between CI and local builds, and between FIPS and non-FIPS pipelines.

## Build Pipeline Architecture

```
spectro-release.yaml (workflow)
│
├── "Build Image" step (non-FIPS)
│   ├── GO_VERSION: 1.24.13
│   └── BUILDER_GOLANG_VERSION: 1.24.13
│
├── "Build Image - FIPS Mode" step
│   ├── GO_VERSION: 1.24.13
│   └── BUILDER_GOLANG_VERSION: 1.24.13
│
└── These env vars override Makefile defaults at CI time
        │
        ▼
    Makefile
    ├── BUILDER_GOLANG_VERSION ?= 1.23        (default, overridden by workflow)
    ├── GO_VERSION ?=1.22.6                    (default, overridden by workflow)
    │
    └── make docker-build-all
            │
            ▼
        Dockerfile
        ├── ARG BUILDER_GOLANG_VERSION
        └── FROM us-docker.pkg.dev/palette-images/build-base-images/golang:${BUILDER_GOLANG_VERSION}-alpine
```

## Current File References

### `.github/workflows/spectro-release.yaml`

The workflow defines two parallel build steps with identical Go versions:

```yaml
      # Non-FIPS build
      -
        name: Build Image
        env:
          REGISTRY: ${{ env.LEGACY_REGISTRY }}
          GO_VERSION: 1.24.13                    # ← Update location 1
          BUILDER_GOLANG_VERSION: 1.24.13        # ← Update location 2
        run: |
          make docker-build-all
          make docker-push-all
      # FIPS build
      -
        name: Build Image - FIPS Mode
        env:
          FIPS_ENABLE: yes
          REGISTRY: ${{ env.FIPS_REGISTRY }}
          GO_VERSION: 1.24.13                    # ← Update location 3
          BUILDER_GOLANG_VERSION: 1.24.13        # ← Update location 4
        run: |
          make docker-build-all
          make docker-push-all
```

### `Makefile`

Default values near the top of the file (used for local development and overridden by the workflow during CI):

```makefile
BUILDER_GOLANG_VERSION ?= 1.23       # ← Update location 5
...
GO_VERSION ?=1.22.6                   # ← Update location 6
```

> **Formatting note**: `GO_VERSION` has no space before `=` in the Makefile. Preserve this exact formatting: `GO_VERSION ?=<version>`.

### `Dockerfile`

The Dockerfile consumes `BUILDER_GOLANG_VERSION` as a build-arg. **This file does not need editing** during a Go upgrade — it dynamically uses whatever version is passed in:

```dockerfile
ARG BUILDER_GOLANG_VERSION
FROM us-docker.pkg.dev/palette-images/build-base-images/golang:${BUILDER_GOLANG_VERSION}-alpine as toolchain
```

## Summary of All 6 Update Locations

| # | File | Variable | Context |
|---|------|----------|---------|
| 1 | `.github/workflows/spectro-release.yaml` | `GO_VERSION` | "Build Image" step (non-FIPS) |
| 2 | `.github/workflows/spectro-release.yaml` | `BUILDER_GOLANG_VERSION` | "Build Image" step (non-FIPS) |
| 3 | `.github/workflows/spectro-release.yaml` | `GO_VERSION` | "Build Image - FIPS Mode" step |
| 4 | `.github/workflows/spectro-release.yaml` | `BUILDER_GOLANG_VERSION` | "Build Image - FIPS Mode" step |
| 5 | `Makefile` | `BUILDER_GOLANG_VERSION` | Default value (line ~15) |
| 6 | `Makefile` | `GO_VERSION` | Default value (line ~24) |

## Step-by-Step Procedure

Given a target Go version (e.g., `1.25.0`):

### Step 1: Update `.github/workflows/spectro-release.yaml`

Replace all 4 version values in the two build steps:

1. In the **"Build Image"** step `env` block, update `GO_VERSION` and `BUILDER_GOLANG_VERSION` to the target version
2. In the **"Build Image - FIPS Mode"** step `env` block, update `GO_VERSION` and `BUILDER_GOLANG_VERSION` to the same target version

Both steps must have identical Go versions.

### Step 2: Update `Makefile`

Replace the 2 default values near the top of the file:

1. Update `BUILDER_GOLANG_VERSION ?= <old>` to `BUILDER_GOLANG_VERSION ?= <new>`
2. Update `GO_VERSION ?=<old>` to `GO_VERSION ?=<new>` (note: no space before `=`)

### Step 3: Verify

1. Read `.github/workflows/spectro-release.yaml` and confirm all 4 version values match the target
2. Read the top of `Makefile` and confirm both default values match the target
3. Confirm no stale version references remain by searching for the old version string across the repository

## Common Pitfalls

### FIPS / Non-FIPS Version Drift

```yaml
# WRONG — versions out of sync between build steps
- name: Build Image
  env:
    GO_VERSION: 1.25.0
    BUILDER_GOLANG_VERSION: 1.25.0

- name: Build Image - FIPS Mode
  env:
    GO_VERSION: 1.24.13              # Still on old version!
    BUILDER_GOLANG_VERSION: 1.24.13
```

Always update both build steps together.

### Makefile Formatting

```makefile
# WRONG — space before = in GO_VERSION
GO_VERSION ?= 1.25.0

# CORRECT — no space before = (matches existing format)
GO_VERSION ?=1.25.0
```

The `BUILDER_GOLANG_VERSION` line does have a space: `BUILDER_GOLANG_VERSION ?= 1.25.0`. Only `GO_VERSION` omits the space.

### Forgetting the Makefile

The Makefile defaults are overridden during CI, so forgetting to update them won't break CI builds. However, **local `make docker-build-all` will use the stale Makefile defaults**, causing confusing local/CI build discrepancies.

### Base Image Availability

Before upgrading, confirm the target version has a corresponding base image at `us-docker.pkg.dev/palette-images/build-base-images/golang:<version>-alpine`. If the image doesn't exist, the Docker build will fail with a pull error.

## Key Files Reference

| File | Purpose |
|------|---------|
| `.github/workflows/spectro-release.yaml` | Release workflow — sets Go versions for FIPS and non-FIPS Docker builds |
| `Makefile` | Build system — default Go versions for local development |
| `Dockerfile` | Multi-stage Docker build — consumes `BUILDER_GOLANG_VERSION` as build-arg |
