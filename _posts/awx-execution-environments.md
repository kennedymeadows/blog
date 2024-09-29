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
---

## Why is there nothing online about this?

I am in the process of building and deploying AWX in an enterprise environment. It's a fantastic tool, especially in the realm of network automation. If you have set out to implement some degree of network automation (NetDevOps, Infrastructure as Code...Net, etc.), you have almost certainly come across Ansible. It's a good tool but there are a lot of issues with integrating it in an enterprise environment:

- There's a learning curve - as nice as the YAML syntax is, it's not the most intuitive especially for network engineers who are used to GUIs.
- It's not very scalable.
- It's difficult to share within a team

AWX solves all of these problems. It's a web-based GUI for Ansible which is deployed in the cloud. Like a lot of great tools, it's open-source and free to use. It seemingly has all of the features of Ansible Automation Platform which, [according to some posts](https://www.reddit.com/r/redhat/comments/w3wo2q/redhat_ansible/) will cost you 10's of thousands of dollars a year for a moderate number of hosts.

I plan to write a longer post about my experiences with deploying AWX which, yes, means learning Kubernetes and spinning up a cluster if you don't already have one available. For now, I want to talk about Execution Environments because there is very little information online about them - even their docs are sparse.

## What are Execution Environments?

When you run Ansible locally, it's simple to use a new collection. Something like...

```bash
$ ansible-galaxy collection install cisco.ios
```

... and you're off. When you run Ansible in AWX, it's more complicated. What AWX does when you trigger a playbook run is to spin up a new container - an Execution Environment - which has all of the collections, python modules, and system-level dependencies which your playbook needs. This is great because it means you can run your playbooks in a consistent environment, regardless of the host you're running them on.

So, to answer a question that I had at one point, "Is there some way to use the `juniper.device` collection without having to create a container image, push it to a container registry, and import it into AWX?" - the answer is no. You need to create an Execution Environment. Hopefully this post will make that process a little easier.

## How do I create an Execution Environment?

The basic workflow is:

1. Create a repository where you define your Execution Environment.
2. Use Ansible-Builder to build a container image from your repository.
3. Push the container image to a container registry - in an enterprise environment, this could be something like Elastic Container Registry (ECR), Azure Container Registry (ACR), or Google Container Registry (GCR).
4. Import the container image into AWX.
5. Use the Execution Environment in your job template.

---

This post is a work in progress. I've spun up a number of these Execution Environments in my enterprise environment. My plan is to spin a few examples up in my personal Github, publish them to Docker Hub, and walk you through how you can do your own. Stay tuned...
