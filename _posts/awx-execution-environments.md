---
title: "AWX Execution Environments"
excerpt: "AWX is a great tool for network automation (and DevOps in general) but there is very little information online about some of the basic building blocks of the tool. This post is an attempt to get you up and running quickly with basic Execution Environments to use in AWX."
coverImage: "/assets/blog/awx-execution-environments/cover.svg"
date: "2024-09-28T05:35:07.322Z"
author:
  name: Simon Lewis
  picture: "/assets/blog/authors/author.svg"
ogImage:
  url: "/assets/blog/awx-execution-environments/cover.svg"
tags: ["Network Automation", "AWX", "Execution Environments", "juniper.device collection"]
---

## Why is there nothing online about this?

**Want to cut to the chase and see the code?** [Here's the Github repository](https://github.com/kennedymeadows/awx-ee-juniper).

I am in the process of building and deploying AWX in an enterprise environment. It's a fantastic tool, especially in the realm of network automation. If you have set out to implement some degree of network automation (NetDevOps, Infrastructure as Code...Net, etc.), you have almost certainly come across Ansible. It's a good tool but there are a lot of issues with integrating it in an enterprise environment:

- There's a learning curve - as nice as the YAML syntax is, it's not the most intuitive especially for network engineers who are used to GUIs.
- It's not very scalable.
- It's difficult to share within a team

AWX solves all of these problems. It's a web-based GUI for Ansible which is deployed in the cloud. Like a lot of great tools, it's open-source and free to use. It seemingly has all of the features of Ansible Automation Platform which, [according to some posts](https://www.reddit.com/r/redhat/comments/w3wo2q/redhat_ansible/) will cost you 10's of thousands of dollars a year for a moderate number of hosts.

I plan to write a longer post about my experiences with deploying AWX which, yes, means learning Kubernetes and spinning up a cluster if you don't already have one available. For now, I want to talk about Execution Environments because there is very little information online about them - even their docs are sparse.

### What are Execution Environments?

When you run Ansible locally, it's simple to use a new collection. Something like...

```bash
$ ansible-galaxy collection install cisco.ios
```

... and you're off. When you run Ansible in AWX, it's more complicated. What AWX does when you trigger a playbook run is to spin up a new container - an Execution Environment - which has all of the collections, python modules, and system-level dependencies which your playbook needs. This is great because it means you can run your playbooks in a consistent environment, regardless of the host you're running them on.

So, to answer a question that I had at one point, "Is there some way to use the `juniper.device` collection without having to create a container image, push it to a container registry, and import it into AWX?" - the answer is no. You need to create an Execution Environment. Hopefully this post will make that process a little easier.

### How do I create an Execution Environment?

The basic workflow is:

1. Create a repository where you define your Execution Environment.
2. Use Ansible-Builder to build a container image from your repository.
3. Push the container image to a container registry - in an enterprise environment, this could be something like Elastic Container Registry (ECR), Azure Container Registry (ACR), or Google Container Registry (GCR).
4. Import the container image into AWX.
5. Use the Execution Environment in your job template.

## An example Execution Environment

This example is going to get you up and running with the `juniper.device` collection. There's a lot you can do with these Execution Environments - you can install system-level dependencies, you can install collections, you can install python modules. This example is going to be very simple - it's just going to install the `juniper.device` collection and the `junos-eznc` python module.

**Want to cut to the chase and see the code?** [Here's the Github repository](https://github.com/kennedymeadows/awx-ee-juniper).

### Step 1: Install Ansible-Builder

You need Ansible Builder version 3.1.0 or later. I tried building this with version 3.0.0 and it didn't work. I believe [this](https://github.com/ansible/ansible-builder/pull/627) was what was causing it to break for me - in any case, version 3.1.0 is working.

```bash
$ pip install ansible-builder
```

Check the version:

```bash
$ ansible-builder --version
```

### Step 2: Create a repository

Create a new repository and add the following files:

```bash
$ touch execution-environment.yml requirements.yml requirements.txt bindep.txt
```

We won't actually use the `bindep.txt` file in this example but I've included it incase you want to build on top of this example.

### Step 3: Define the Execution Environment

Edit the `execution-environment.yml` file:

```yaml
---
version: 3.1

dependencies:
  ansible_core:
    package_pip: ansible-core
  ansible_runner:
    package_pip: ansible-runner
  galaxy: requirements.yml
  python: requirements.txt
  system: bindep.txt

images:
  base_image:
    name: quay.io/ansible/awx-ee:latest

additional_build_steps:
  append_base:
    - RUN $PYCMD -m pip install -U pip

```

The good thing is that the `awx-ee` base image has taken care of a lot of the dependencies for us. We just need to install the `juniper.device` collection and the `junos-eznc` python module.

### Step 4: Define the dependencies

Edit the `requirements.yml` file:

```yaml
---
collections:
  - name: juniper.device

# roles goes here if you are going to add some later

```

Edit the `requirements.txt` file:

```txt
junos-eznc
```

### Step 5: Build the container image

```bash
ansible-builder build -v3.1 -t awx-ee-juniper:latest --container-runtime=docker
```

This is presuming that you have Docker installed and running. If you're using Podman, you can drop the whole `--container-runtime=docker` flag - Podman is the default.

### Step 6: Push the container image

I'm going to let you figure this one out. In the workplace you most likely have a private container registry - you'll need to push the image to that registry.

### Step 7: Import the container image into AWX

In the AWX GUI, go to `Settings` -> `Execution Environments` -> `Add Execution Environment`. Fill in the details and click `Save`.

You can also use this example now in your environment - it's available [on Docker Hub](https://hub.docker.com/r/levtolstoi/awx-ee-juniper). Within the AWX console, you can import it by specifying the image name `levtolstoi/awx-ee-juniper:latest`.

## Conclusion

That's it! You've created an Execution Environment. You can now use this Execution Environment in your job templates. You can also build on this example - you can install system-level dependencies, you can install collections, you can install python modules. You can also use this example to create Execution Environments for other collections.

Here are some resources if you want to learn more:

- [Ansible for Junos OS](https://www.juniper.net/documentation/us/en/software/junos-ansible/ansible/topics/concept/junos-ansible-modules-overview.html)
- [AWX Execution Environments](https://ansible.readthedocs.io/projects/awx/en/latest/userguide/execution_environments.html)
- [Ansible Builder Execution Environment documentation](https://ansible.readthedocs.io/projects/builder/en/latest/definition/)

