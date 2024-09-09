<a name="readme-top"></a>

<div align="center">
    <img src="https://i.imgur.com/qmPQWGj.png" alt="Logo" width="100" height="100">

<h3 align="center">nekoya</h3>

  <p align="center">
    discord bot for cat lovers
    <br />
    <br />
    <a href="https://github.com/riceandbeas/nekoya/issues/new?labels=bug&template=bug_report.md">Report Bug</a>
    ·
    <a href="https://github.com/riceandbeas/nekoya/issues/new?labels=enhancement&template=feature_request.md">Request Feature</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
        <ul>
          <li><a href="#features">Features</a></li>
          <li><a href="#etymology">Etymology</a></li>
        </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
  </ol>
</details>

## About The Project
Nekoya is a fun and interactive Discord bot made for cat enthusiasts.

### Features
- **Slash Commands:**
  - `/fact`: Replies with a random cat fact.
  - `/pic [breed]`: Replies with a random cat picture. Optionally, a breed can be specified.
  - `/http [status_code]`: Replies with an HTTP cat image corresponding to the given status code.

- **Text Handlers:**
  - Responds with a cat's meow tailored to the language of the question, such as "what do cats say?".
    <details>
      <summary>Supported Languages</summary>
      <ul>
      <li>English: "what do cats say?" (meow)</li>
      <li>Portuguese: "o que gatos dizem?" (miau)</li>
      </ul>
    </details>

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Etymology
The name **Nekoya** is inspired by two sources:

1. Neko (猫) – The Japanese word for "cat", reflecting the bot's core theme of cats.
2. Nekkoya (내꺼야 (Pick Me)) – A reference to the theme song of the competition show Produce 48.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Usage

### 1. Create a Discord Application
Set up your Discord bot by creating a new bot on the [Discord Developer Portal](https://discord.com/developers/applications) and invite it to your server.

If not familiar with this process, take a look at the [official documentation](https://discord.com/developers/docs/quick-start/getting-started).

Recommended permissions: `Send Messages`, `Read Message History`, `Use Slash Commands`.

### 2. Clone the Repository
If you haven't already, clone the Nekoya repository to your local machine:
```bash
git clone https://github.com/riceandbeas/nekoya.git
cd nekoya
```

### 3. Set Up Environment Variables
1. Create a `.env` file in the project’s root directory based on the provided [`.env.example`](/.env.example) file:
```bash
cp .env.example .env
```

2. Fill in the necessary values in the `.env` file.

### 4. Install Dependencies
Ensure all Go module dependencies are downloaded and installed:
```bash
go mod download
```

### 5. Run the Bot
To start the bot, run the following command:
```bash
go run main.go start
```

### 6. Interact with Nekoya
Once the bot is running, it will begin listening for interactions in your Discord server.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## License

Distributed under the GNU Affero General Public License. See `LICENSE` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>
