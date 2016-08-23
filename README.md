# noto-emoji-iconset

Iconset for [`iron-icon`](https://elements.polymer-project.org/elements/iron-icon) to use Google's Emojis.

## How to install
You can clone this repo directly to your server but I recommend using bower:

`bower install --save raulsntos/noto-emoji-iconset`

## How to use the set
To use the set simply import this set and use it like any other iconset. Use the prefix **emoji** followed by colon (**:**) and the emoji in unicode (`🎉`). Example:
```html
<html>
  <head>
    ...
    <link rel="import" href="bower_components/noto-emoji-iconset/noto-emoji-iconset.html">
    ...
  </head>
  <body>
    ...
      <iron-icon icon="emoji:🎉"></iron-icon>
    ...
  </body>
</html>
```

## But it's hard to type emojis in my laptop! :angry:
I know! That's why there's also a Polymer element included `emoji-icon` which lets you use the emoji shortname (like you do in GitHub). Example:
```html
<html>
  <head>
    ...
    <link rel="import" href="bower_components/noto-emoji-iconset/emoji-icon.html">
    ...
  </head>
  <body>
    ...
      <emoji-icon emoji="tada"></emoji-icon>
    ...
  </body>
</html>
```

### Nice features of `emoji-icon`
- You can use shortnames instead of typing the emoji (making it easier to use when you are not developing using your phone :wink:)
- The element uses a dictionary to translate emoji shortnames to unicode, the dictionary is stored in `emoji.json`. The element loads the JSON file **only once** and waits until the Promise is resolved to load the emoji.

## Can I change the size?
Yes! If you are using `iron-icon`, see how to do it in their [documentation](https://elements.polymer-project.org/elements/iron-icon#styling). If you are using `emoji-icon` you can also find how in `iron-icon`'s documentation because it's the exactly the same but replacing `iron` with `emoji` in the CSS variables.

## How to build the iconset
You can build the iconset yourself by using the build.go file included in this repository, simple use `go run build.go`. Note that you need to have `git` installed since the script will clone the Noto GitHub repository (you can also download the repository manually, the script will not clone it as long as a folder named `noto-emoji` exists in the root of this project). If you have already cloned or downloaded the Noto repository but want to update it to the latest version use the flag `-update` when running the script: `go run build.go -update`. The script will replace the `noto-emoji-iconset.html` file.

## Known issues
- I'm using the SVG icons provided by Google in the [Noto repository](https://github.com/googlei18n/noto-emoji) and they are currently outdated so there are a few missing emojis, when Google updates their repository I'll include the new emoji. If the emoji icon is not found. If you want to know the state of this issue check the Noto repository issue [#62](https://github.com/googlei18n/noto-emoji/issues/62).
