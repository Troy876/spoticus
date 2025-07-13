# Spoticus

**Spoticus** is a Slack-integrated bot that provisions Kubernetes and OpenShift clusters on AWS using **spot instances** — offering a cost-effective and on-demand environment for development and testing.

---

## 🔍 Overview

Spoticus listens to Slack commands and interacts with the cloud to:

- 🟢 Launch upstream Kubernetes or OpenShift clusters
- ⚡ Provision them using AWS **spot instances** for optimal cost savings with mapt-operator
- 📏 Support various compute tiers (`medium`, `large`, `xlarge`)
- 📡 Provide instant feedback and status via Slack messages

---

## 🚀 Supported Commands

Spoticus currently supports the following command from Slack:

### `launch`

Provision a new cluster based on your desired platform and size.

#### Syntax

```bash
launch <cluster_type> <size>
```

#### Supported Cluster Types

- `k8s` — Standard Kubernetes
- `openshift` — Red Hat OpenShift

#### Supported Sizes

| Size    | CPUs     | RAM       |
|---------|----------|-----------|
| medium  | 8        | 32 GB     |
| large   | 16       | 64 GB     |
| xlarge  | 32       | 128 GB    |

#### Slack commands events

``` bash
launch k8s large
launch openshift medium
```

> All clusters are created using AWS **spot instances** to ensure maximum efficiency and reduced cloud spend.

---

## 🛠️ Getting Started

### Requirements

- Go 1.23+
- Slack Bot Token and App configuration

### Build the Project

```bash
make build
```

Binary will be output to: `bin/spoticus`

### Run Locally

```bash
make run
```

> Make sure your Slack bot token and mapt-operator are set in your environment or configuration.

---

## 🧪 Development

### Format & Lint

```bash
make fmt
make lint
```

### Clean Artifacts

```bash
make clean
```

---

## 📂 Project Structure

```text
cmd/
  spoticus/           # Main entrypoint for the bot
internal/
  slack/              # Slack command handling
bin/                  # Compiled binaries (ignored in Git)
```
