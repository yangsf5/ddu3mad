import type { PlasmoCSConfig } from "plasmo"

import { quote } from "~core/quote"

export const config: PlasmoCSConfig = {
  matches: ["https://www.baidu.com/*"]
}

window.addEventListener("load", () => {
  alert(quote)
})