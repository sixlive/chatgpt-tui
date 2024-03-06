# ChatGPT tui README

My first ever terminal UI! Everything is stored locally on sqlite and written in Go!

## Technologies
- SQLite
- [bubbletea](https://github.com/charmbracelet/bubbletea)
- Go

## Installation

Please make sure that you expose a `CHAT_GPT_API_KEY` inside of your environment; we require it to make api calls!

Set up your [api key](https://platform.openai.com/api-keys)

```bash
export CHAT_GPT_API_KEY="some-key" # you would want to export this in your .zshrc
brew tap tearingitup786/tearingitup786
brew install chatgpt-tui
chatgpt-tui
```

![tui demo](./tui-demo.gif)

## Global Keybindings

- `Tab`: Change focus between panes. The currently focused pane will be highlighted with a pink border.
- `Ctrl+o`: Toggles zen mode

## Chat Messages Pane

- `y`: Copies the last message from ChatGPT into your clipboard.
- `Y`: Copies all messages from the ChatGPT session into your clipboard.

## Settings Pane

- `m`: Opens an input dialog to change the model.
- `f`: Opens an input dialog to change the frequency of updates.
- `t`: Opens an input dialog to set the maximum number of tokens per message.

## Sessions Pane

- `Ctrl+N`: Creates a new session.
- `d`: Deletes the currently selected session from the list.
- `Enter`: Switches to the session that is currently selected.

Please refer to this guide as you navigate the TUI. Happy exploring!
