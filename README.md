<div id="top"></div>

<!-- PROJECT SHIELDS -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links-->
<div align="center">

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![Wiki][wiki-shield]][wiki-url]

</div>

<!-- PROJECT LOGO -->
<br />
<!-- UPDATE -->
<div align="center">
  <a href="https://github.com/metakgp/mfins">
     <img width="140" alt="image" src="https://raw.githubusercontent.com/metakgp/design/main/logos/black-large.jpg">
  </a>

  <h3 align="center">MFINS</h3>

  <p align="center">
  <!-- UPDATE -->
    <i>My Freakin' Internal Noticeboard Scrapper</i>
    <br />
    <a href="https://UPDATE.metakgp.org">Website</a>
    ·
    <a href="https://github.com/metakgp/mfins/issues">Request Feature / Report Bug</a>
  </p>
</div>


<!-- TABLE OF CONTENTS -->
<details>
<summary>Table of Contents</summary>

- [About The Project](#about-the-project)
  - [Supports](#supports)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Contact](#contact)
  - [Maintainer(s)](#maintainers)
  - [creators(s)](#creators)
- [Additional documentation](#additional-documentation)

</details>


<!-- ABOUT THE PROJECT -->
## About The Project
<!-- UPDATE -->
<div align="center">
  <a href="https://github.com/metakgp/mfins">
    <img width="80%" alt="image" src="https://user-images.githubusercontent.com/86282911/206632547-a3b34b47-e7ae-4186-a1e6-ecda7ddb38e6.png">
  </a>
</div>

This a microservice designed to send notceboard messages for the internal noticeboard to Naarad (NTFY). It uses the SSO token and jsessionid to access the noticeboard present in the erp. It checks for a new message every 2 minutes.
<p align="right">(<a href="#top">back to top</a>)</p>

<div id="supports"></div>

<!-- ### Supports: -->

<!-- <p align="right">(<a href="#top">back to top</a>)</p> -->

## Getting Started

To set up a local instance of the application, follow the steps below.

### Prerequisites
The following dependencies are required to be installed for the project to function properly:
<!-- UPDATE -->
* [golang](https://go.dev/doc/install)

<p align="right">(<a href="#top">back to top</a>)</p>

### Installation

_Now that the environment has been set up and configured to properly compile and run the project, the next step is to install and configure the project locally on your system._
<!-- UPDATE -->
1. Clone the repository
   ```sh
   git clone https://github.com/metakgp/mfins.git
   ```
2. Change directory to the folder
   ```sh
   cd ./mfins
   ```
3. Copy the `.env.example` to `.env` file
    ```shell
    cp .env.example .env
    ```
4. Fill the .env files with your erp credentials

5. Create `lastmsg.json` and add `{}` in it
    ```sh
      touch lastmsg.json && echo {} > lastmsg.json
    ```
6. Create `security_question.json` and add `{}` to it
    ```sh
      touch security_question.json && echo {} > security_question.json
    ```

7. Create a google OAuth client secret and client id and add it to `client_secret.json`

8. Execute the script
   ```sh
    go run main.go
   ```

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage
<!-- UPDATE -->
Use this space to show useful examples of how this project can be used. Additional screenshots, code examples and demos work well in this space.

<div align="center">
  <a href="https://github.com/metakgp/mfins">
    <img width="80%" alt="image" src="https://user-images.githubusercontent.com/86282911/206632640-40dc440e-5ef3-4893-be48-618f2bd85f37.png">
  </a>
</div>

<p align="right">(<a href="#top">back to top</a>)</p>

## Contact

<p>
  📫 Metakgp -
  <a href="https://slack.metakgp.org">
    <img align="center" alt="Metakgp's slack invite" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/slack.svg" />
  </a>
  <a href="mailto:metakgp@gmail.com">
    <img align="center" alt="Metakgp's email " width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/gmail.svg" />
  </a>
  <a href="https://www.facebook.com/metakgp">
    <img align="center" alt="metakgp's Facebook" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/facebook.svg" />
  </a>
  <a href="https://www.linkedin.com/company/metakgp-org/">
    <img align="center" alt="metakgp's LinkedIn" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/linkedin.svg" />
  </a>
  <a href="https://twitter.com/metakgp">
    <img align="center" alt="metakgp's Twitter " width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/twitter.svg" />
  </a>
  <a href="https://www.instagram.com/metakgp_/">
    <img align="center" alt="metakgp's Instagram" width="22px" src="https://raw.githubusercontent.com/edent/SuperTinyIcons/master/images/svg/instagram.svg" />
  </a>
</p>

### Maintainer(s)

The currently active maintainer(s) of this project.

<!-- UPDATE -->
- [Bikram Ghuku](https://github.com/bikram-ghuku)

### Creator(s)

Honoring the original creator(s) and ideator(s) of this project.

<!-- UPDATE -->
- [Arpit Bhardwaj](https://github.com/proffapt)

<p align="right">(<a href="#top">back to top</a>)</p>

## Additional documentation

  - [License](/LICENSE)
  - [Code of Conduct](/.github/CODE_OF_CONDUCT.md)
  - [Security Policy](/.github/SECURITY.md)
  - [Contribution Guidelines](/.github/CONTRIBUTING.md)

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->

[contributors-shield]: https://img.shields.io/github/contributors/metakgp/mfins.svg?style=for-the-badge
[contributors-url]: https://github.com/metakgp/mfins/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/metakgp/mfins.svg?style=for-the-badge
[forks-url]: https://github.com/metakgp/PROJECT_NAME/network/members
[stars-shield]: https://img.shields.io/github/stars/metakgp/mfins.svg?style=for-the-badge
[stars-url]: https://github.com/metakgp/mfins/stargazers
[issues-shield]: https://img.shields.io/github/issues/metakgp/mfins.svg?style=for-the-badge
[issues-url]: https://github.com/metakgp/mfins/issues
[license-shield]: https://img.shields.io/github/license/metakgp/mfins.svg?style=for-the-badge
[license-url]: https://github.com/metakgp/mfins/blob/master/LICENSE
[wiki-shield]: https://custom-icon-badges.demolab.com/badge/metakgp_wiki-grey?logo=metakgp_logo&style=for-the-badge
[wiki-url]: https://wiki.metakgp.org
[slack-url]: https://slack.metakgp.org
