# Helm Operator

> **🛑 Upgrade Advisory**
>
> This documentation is for Helm Operator (v1) which has [reached its end-of-life in November 2022](https://fluxcd.io/blog/2022/10/september-2022-update/#flux-legacy-v1-retirement-plan).
>
> We strongly recommend you familiarise yourself with the newest Flux and [migrate as soon as possible](https://fluxcd.io/flux/migration/).
>
> For documentation regarding the latest Flux, please refer to [this section](https://fluxcd.io/flux/).

The Helm Operator is a [Kubernetes operator](https://kubernetes.io/docs/concepts/extend-kubernetes/operator/),
allowing one to declaratively manage Helm chart releases. Combined with
[Flux](https://github.com/fluxcd/flux) this can be utilized to automate
releases in a GitOps manner, but the usage of Flux is not a strict
requirement.

The desired state of a Helm release is described through a Kubernetes
Custom Resource named `HelmRelease`. Based on the creation, mutation or
removal of a `HelmRelease` resource in the cluster, Helm actions are
performed by the operator.

![Helm Operator](../docs/_files/fluxcd-helm-operator-diagram.png)

## Features

* Declarative install, upgrade, and delete of Helm releases
* Pulls chart from _any_ chart source;
    * Public or private Helm repositories over HTTP/S
    * Public or private Git repositories over HTTPS or SSH
    * Any other public or private chart source using one of the available
      [Helm downloader plugins](https://helm.sh/docs/topics/plugins/#downloader-plugins)
* Allows Helm values to be specified;
    * In-line in the `HelmRelease` resource
    * In (external) sources, e.g. `ConfigMap` and `Secret` resources,
      or a (local) URL
* Automated purging on release install failures
* Automated (optional) rollback on upgrade failures
* Automated (optional) helm test gating of installs and upgrades.
* Automated image upgrades [using Flux](https://github.com/fluxcd/flux/blob/master/docs/references/helm-operator-integration.md)
* Automated (configurable) chart dependency updates for Helm charts
  from Git sources on install or upgrade
* Detection and recovery from Helm storage mutations
  (e.g. a manual Helm release that was made but conflicts with the
  declared configuration for the release)
* Parallel and scalable processing of different `HelmRelease` resources
  using workers
* Supports both Helm 2 and 3
* Supports automatic Helm migrations from v2 to v3

## Get started

Get started installing the Helm Operator:

- [Quickstart](get-started/quickstart.md)
- [Get started using Helm](get-started/using-helm.md)
- [Get started using Kustomize](get-started/using-kustomize.md)
- [Get started using YAMLs](get-started/using-yamls.md)

## Integrations

As the Helm Operator is Open Source, integrations are very straight-forward. Here are
a few popular ones you might want to check out:

- [Progressive Delivery workshop](https://helm.workshop.flagger.dev/)
- [Managing Helm releases the GitOps way](https://github.com/fluxcd/helm-operator-get-started)
- [GitOps for Istio Canary deployments](https://github.com/stefanprodan/gitops-istio)

## Getting help

If you have any questions about the Helm Operator and continuous delivery:

- Invite yourself to the <a href="https://slack.cncf.io" target="_blank">CNCF community</a>
  slack and ask a question on the [#flux](https://cloud-native.slack.com/messages/flux/)
  channel.
- To be part of the conversation about Helm Operator's development, join the
  [flux-dev mailing list](https://lists.cncf.io/g/cncf-flux-dev).
- [File an issue.](https://github.com/fluxcd/helm-operator/issues/new)

Your feedback is always welcome!
