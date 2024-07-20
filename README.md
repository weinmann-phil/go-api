<!-- DISCLAIMER -->
<!-- This README uses the template provided by 
*** [othneildrev](https://github.com/othneildrew/Best-README-Template/blob/master/README.md)
*** and is licensed under the MIT creative commons license. (2022-09-30)
*** Please support the channel.
-->
# Go API
<!-- <a name="readme-top"></a> -->

<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/weinmann-phil/go-api">
    <img src="img/Go-Logo_Blue.png" alt="Logo" width="80" height="80">
  </a>

  <h3 align="center">Simple Go API</h3>

  <p align="center">
    Simple Go API 
    <br />
    <a href="https://github.com/weinmann-phil/go-api"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/weinmann-phil/go-api">View Demo</a>
    ·
    <a href="https://github.com/weinmann-phil/go-api/issues">Report Bug</a>
    ·
    <a href="https://github.com/weinmann-phil/go-api/issues">Request Feature</a>
  </p>
</div>

<!-- OVERVIEW -->
## 0. Overview

This project falls under the category of `learning` and is for private purposes only.

The focus of this repo is to write backend logic for a RESTful API, focusing on GORM and Gin-Gonic.
In addition to that, this repo also is using `Podman` as replacement for `Docker`.


Here's why:

* Go is awesome 😻
* I wanted to learn more about the inner workings of `gin-gonic` 🍸
* I wanted to learn more about the inner workings of `GORM` 😺
* I wanted to learn more about `podman` 🐳

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol start="0">
    <li>
      <a href="#0.-overview">Overview</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
  </ol>
</details>

<p align="right">(<a href="#go-api">back to top</a>)</p>

<!-- Dependencies/technologies -->
### Built With

This project we are using Go version `go1.22.4`, developed entirely on `linux/amd64` system.
<!-- Please check out their respective documentation: -->

<!-- [![Terraform][Terraform]][Terraform-url] -->

<!-- [![GitLab][GitLab]][GitLab-url] -->

<p align="right">(<a href="#go-api">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

For a local setup, you need to have `Docker` and `Terraform` installed and have
an official Docker container for the GitLab community edition running.
If you have not yet installed these items, please take the following steps:

### Prerequisites

* Install Docker

  ```sh
  sudo apt install gnome-terminal
  sudo apt update
  sudo apt install ./docker-desktop-<version>-<arch>.deb
  ```

* Run an instance of PostgreSQL

  ```sh
  docker run --name postgres \
    -e POSTGRES_PASSWORD=$PASSWORD \
    -e POSTGRES_USER=$USER \ 
    -e POSTGRES_DB=$DATABASE_NAME \ 
    -p 5432:5432 \ 
    -d postgres
  ```

### Installation

<!-- > __NOTE__:
>
> This is a sample usage of this project.
> If you are applying this within any environments other than a local test environment,
> please mind to change the settings for the provider configuration and the backend
> configurations.
>
> Otherwise, this will not work.

1. Set up your self-hosted GitLab system

   ```sh
   sudo docker run -d \
   -p 443:443 -p 80:80 -p 22:22 \
   --hostname localhost \
   --name gitlab-ce \
   --restart always \
   -v $GITLAB_HOME/config:/etc/gitlab \
   -v $GITLAB_HOME/logs:/var/log/gitlab \
   -v $GITLAB_HOME/data:/var/opt/gitlab \
   --shm-size 256m \
   gitlab/gitlab-ce:latest
   ```

1. Create an access token with administrative privileges
   
   <details>
     <summary>Create access token</summary>

     ![create-access-token](./img/gitlab_access-token.png)

   </details>

1. Enter token and a list of users into terraform.tfvars

1. Switch directory to the workspace

   ```sh
   cd environments/gitlab/
   ```

1. Initialize project

   ```sh
   terraform init
   ```

1. Apply changes to your GitLab

   ```sh
   terraform apply
   ```

<p align="right">(<a href="#go-api">back to top</a>)</p> -->


<!-- USAGE EXAMPLES -->
## Usage

There is no particular usage yet.

<p align="right">(<a href="#go-api">back to top</a>)</p>


<!-- ROADMAP -->
## Roadmap

- [x] Add License
- [x] Add Readme
- [x] Extend logic to include all CRUD operations
- [x] Containerize the entire logic with Podman
- [ ] Set up a compose project to run project with all its dependencies with Podman compose
   
See the [open issues](https://github.com/weinmann-phil/go-api/issues) for a full list of proposed features (and known issues).

<p align="right">(<a href="#go-api">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

This is my first project in an attempt in giving back to the community.
Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. 
You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
1. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
1. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
1. Push to the Branch (`git push origin feature/AmazingFeature`)
1. Open a Pull Request

<p align="right">(<a href="#go-api">back to top</a>)</p>


<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE` for more information.

<p align="right">(<a href="#go-api">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Philip Weinmann - Philip.Weinmann@protonmail.com

Project Link: [https://github.com/weinmann-phil/go-api](https://github.com/weinmann-phil/go-api)

<p align="right">(<a href="#go-api">back to top</a>)</p>


<!-- ACKNOWLEDGMENTS -->
## Acknowledgments

I am grateful to the people that provided the following resources, which I have 
used to create this project.

* [README template](https://github.com/othneildrew/Best-README-Template)
* [GitHub - withsilasogar - REST API Tutorial with Go](https://github.com/withsilasogar/rest-api-tutorial-with-go/tree/main)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/weinmann-phil/go-api.svg?style=for-the-badge
[contributors-url]: https://github.com/weinmann-phil/go-api/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/weinmann-phil/go-api.svg?style=for-the-badge
[forks-url]: https://github.com/weinmann-phil/go-api/network/members
[stars-shield]: https://img.shields.io/github/stars/weinmann-phil/go-api.svg?style=for-the-badge
[stars-url]: https://github.com/weinmann-phil/go-api/stargazers
[issues-shield]: https://img.shields.io/github/issues/weinmann-phil/go-api.svg?style=for-the-badge
[issues-url]: https://github.com/weinmann-phil/go-api/issues
[license-shield]: https://img.shields.io/github/license/weinmann-phil/go-api.svg?style=for-the-badge
[license-url]: https://github.com/weinmann-phil/go-api/blob/main/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://linkedin.com/in/philipweinmann
[product-screenshot]: img/Go-Logo_Blue.png
<!-- [Terraform]: https://img.shields.io/badge/terraform-4A235A?style=for-the-badge&logo=terraform -->


## Referecenes

* [Go - Gin]()
* [Go - Gorm]()