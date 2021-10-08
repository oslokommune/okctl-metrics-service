
# Contributing

## Intro

First of all; thank you for the interest!

## CD

Any merge to master prompts a release. Any release prompts a deploy. 

In a more systematic description:
1. Commit to master
2. `.github/workflows/cd.yml` triggers which build and pushes an image to `ghcr.io` and notifies the infrastructure repo
    [okctl-infrastructure](https://github.com/oslokommune/okctl-infrastructure)
3. An action in the infrastructure repo does the necessary changes to infrastructure manifests
4. ArgoCD picks up the change in the infrastructure repo and deploys them to the cluster
