# Introduction to Containers
[`@damdo`](https://twitter.com/damdo)
<!-- .slide: data-transition="none" -->

---
<!-- .slide: data-transition="none" -->

## Whoami
```
$ whoami
total 3
-rw-r--r--  name: dam (Damiano Donati)
-rw-r--r--  work: Software Engineer @ Pusher (pusher.com)
-rw-r--r--  twitter: @damdo
-rw-r--r--  github: @damdo
```
Note:
- d
- d

---
<!-- .slide: data-transition="none" -->
## Disclaimer

This is a beginners/introductory workshop

---

<!-- .slide: data-transition="none"; -->

## what is a container
<div style="font-size:30px">

In generic terms:

```text
con·tain·er | kən-ˈtā-nər

"A container is any receptacle or enclosure
for holding a product and is used to store, 
package, and ship it"
```
---
<!-- .slide: data-transition="none"; -->
## what is a container
<div style="font-size:30px">

- We can _contain_
    - _food_: kitchen containers 🥫
    - _goods_: ship containers 🚢
    - _software_: "software" containers 📦
    - and more ...

- We will of course focus on Software containers

---
## containers in software
<div style="font-size:30px">

<!-- .slide: data-transition="none"; -->
Various things that are referred as a Container in Software:
- type container: abstract data type

- file container: stores data together (i.e. mkv)

- app container: an app abstraction/isolation method

We will focus on the latter, and from now on we will refer to it simply as "container".

---
<!-- .slide: data-transition="none"; -->
## A bit of history
<div style="font-size:30px">

To understand what a container is and why it has been concieved we need to look a bit back 🔙 to a recent past... 
---
<!-- .slide: data-transition="none" -->
## Before server virualization era
<div style="font-size:30px">

- **OS** always directly installed on **bare-metal** machines

- setup of machines, tedious/expensive in time & money

- replacement of machines: highly demanding task 🙇

---
<!-- .slide: data-transition="none" -->
## servers updates
<div style="font-size:30px">

Servers were updated and modified in-place:
- software packages
- configurations
- files 
- deployments

were changed through manual intervention (SSH)<br> or via automated scripts.

---
<!-- .slide: data-transition="none" -->
## virtualization came along
<div style="font-size:30px">

- easier to setup machines
- easier to snapshot/clone machines

- **machines** were still updated **in place** (SSH, scripts)
- **machines** where still kept **running** 🏃‍♀️ as long as possible
- machine 🆙 **uptime** ⏱️ was considered a great achievement

---
<!-- .slide: data-transition="none" -->
## machines == unique systems
<div style="font-size:30px">

- machines being **hard to replace** fueled a *vicuous circle*.

- machines became **unique systems** constantly receiving software updates and patches to stay afloat

- people were **scared to loose them**, after all updates and patches, **impossible to replicate their state from zero**

---
<!-- .slide: data-transition="none" -->
## The Pet vs Cattle tale
<div style="font-size:30px">

Randy Bias, at CloudConnect in 2012,
highlights this issue with a similitude:
```text
“ [...] People give names to pets, they are unique, 
they love them and they are lovingly 
hand raised with care,
like for example when they are sick, 
people nurse them until they heal. 
You can feed them and they grow in size, 
when they leave, people misses them.”
```
```text
“Cattle instead are numbered, they 
look all the same and if they are 
ill people can replace them with 
other cattle that does the same.

The Pet model is what the current 
(server) infrastructures resemble and instead 
the Cattle model is what future 
(server) infrastructures should aim to be”.
```
---
<!-- .slide: data-transition="none" -->
## A new idea of server
<div style="font-size:30px">

So there is a new portraied idea of infrastructure<br>
as one composed by servers that are:
- **ephemeral** - (can last a very short time)
- **immutable** - (don't change during their lifetime)
- **disposable** - (used once and then thrown away)

---
<!-- .slide: data-transition="none" -->
## A new idea of server
<div style="font-size:30px">
But such **flexibility was hard** to achieve **with VMs** due to:
<br><br>

<ul style="font-size:30px">
<li>boot times</li>
<li>full stack of software to bootstrap (kernel, libs, apps)</li>
</ul>
<br><br>
In order to gain more flexibility there was a need to<br> **decouple the apps** from the underlying **stack**.
</div>
---
<!-- .slide: data-transition="none" -->
## The advent of Linux containers
<p style="font-size:25px">
Amongst the first to realize this there was Google, which in 2006 proposed a patch to the Linux Kernel to introduce "Containers" (merged in 2008 as **cgroups**).
</p>

<img src="/images/google-linux.png" width="600px">

---
<!-- .slide: data-transition="none" -->
## The advent of Linux containers
<p style="font-size:25px">
And together with other concepts and ideas contributed to lay the foundation of Linux containers:
</p>

<img src="/images/linux-containers-timeline.jpeg" width="900px">

<p style="font-size:25px">
Credits: <br>
https://twitter.com/DanielKrook
https://twitter.com/bibryam/status/885221134408527872/photo/1

---
<!-- .slide: data-transition="none" -->
## Containers Definition
<p style="font-size:25px">
After several iterations over these ideas in 2013, Docker emerged, bringing to the masses containers as we know them now:

```text
A container is defined by two components:

- a container image 📦️:
a lightweight, standalone and executable 
package of software that includes
everything needed to run an application: 
app code / binaries,
system tools, libraries and settings.

- a container (runtime) ⚙️:
a self contained and portable group of 
one or more running processes,
packaged with all the code and dependencies
so the application can be quickly and reliably
ran over several computing environments. 
```
---
<!-- .slide: data-transition="none" -->
## VMs vs Containers

<img src="/images/vm-vs-container.png" width="700px">

<p style="font-size:20px">
Credits:<br>
Joe Yankel, Carnagie Mellon<br>
https://www.microcontrollertips.com/containerization-differs-virtual-machines-faq/

Note:
here on the left we have an example of 3 VMs,
they are sitting on a bare metal server, with their own OS, like Ubuntu or Windows Server, and on top of it we have an Hypervisor like VMWare, Citrix or KVM.

The 3 VMs are in fully emulated machine environment and in turn they have their own operating system (that's N + 1 OSes on one single machine!!). Then eachone of them has their bin and libs + the apps. That's a lot of overhead.

on the left we have a container setup.
We still have a host machine with an host os (this must be one that supports Docker) and then on top of it we have all the bins and libs (and distribution) that we need in order to run our app. That's way less overhead.

---
<!-- .slide: data-transition="none" -->
## containers advantages
<ul style="font-size:30px">
<li>resource efficiency and density</li>
<li>platform independence: build once, run anywhere</li>
<li>isolation and resource sharing</li>
<li>speed: start, create, destroy in seconds</li>
<li>increased dev productivity and faster pipelines</li>
</ul>

---
<!-- .slide: data-transition="none" -->
## Docker
<div style="font-size:30px">

Docker allows us to define, build and run containers quickly.<br>
Let's play with it.<br><br>
visit: [tinyurl.com/intro-docker](https://tinyurl.com/intro-docker)

---
<!-- .slide: data-transition="none" -->
## Fine
