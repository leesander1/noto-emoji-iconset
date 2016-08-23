# noto-emoji-iconset

Iconset for [`iron-icon`](https://elements.polymer-project.org/elements/iron-icon) to use Google's Emojis. [See the Documentation](https://raulsntos.github.io/noto-emoji-iconset)

[See for yourself](https://raulsntos.github.io/noto-emoji-iconset/components/noto-emoji-iconset/demo).

![Emojis](https://github.com/raulsntos/noto-emoji-iconset/raw/master/hero.png)

## How to install
You can clone this repo directly to your server but I recommend using bower:

`bower install --save raulsntos/noto-emoji-iconset`

## Important note
The emojis will not render properly if you don't use Shadow DOM so be sure to enable it.

**The good news:** Polymer will switch to Shadow DOM by default eventually (soon).

Take a look at the [demo](https://github.com/raulsntos/noto-emoji-iconset/blob/master/demo/index.html) to see how I'm using the iconset. Just be sure to add this code before loading Polymer (or every element implemented with Polymer):
```html
<script>
window.Polymer = {
  dom: 'shadow'
}
</script>
```

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
- I would prefer to make `emoji-icon` extend from `iron-icon` instead of creating a new element with an `iron-icon` tag inside and redefining every CSS variable but Polymer does not seem to support extending custom elements yet. See issue [#2280](https://github.com/Polymer/polymer/issues/2280) in the Polymer repo.
