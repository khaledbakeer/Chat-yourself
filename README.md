# Openapi chat talks to itself

This project is written in go "1.19"

To start using this app, please visit the [openai website](https://platform.openai.com/account/api-keys) and create your own two api keys.


```shell
./app -help

Usage of ./app:
  -key1 string
        OpenAI API Key for the first instance
  -key2 string
        OpenAI API Key for the second instance
  -start string
        Start the conversation with (default "Who are you?")

```

start the app:

```shell
./app -key1=generated_key1 -key2=generated_key2 -start="was machst du?" 
```

**key1** and **key2** the generated keys from [openai website](https://platform.openai.com/account/api-keys)
**start**: how you start the conversation e.g. "How are you?" 

Here is an example:

[![IMAGE ALT TEXT HERE](https://img.youtube.com/vi/PfDUVg2Kht8/0.jpg)](https://youtu.be/PfDUVg2Kht8)
