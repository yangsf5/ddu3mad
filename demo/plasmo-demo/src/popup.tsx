import Button from "antd/es/button"
import { useEffect, useRef, useState } from "react"

import { ThemeProvider } from "~theme"

function IndexPopup() {
  const [data, setData] = useState("")

  const iframeRef = useRef<HTMLIFrameElement>(null)
 
  useEffect(() => {
    window.addEventListener("message", (event) => {
      console.log("EVAL output: " + event.data)
    })
  }, [])

  return (
    <ThemeProvider>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          padding: 16
        }}>
        <h2>
          Welcome to your{" "}
          <a href="https://www.plasmo.com" target="_blank">
            Plasmo
          </a>{" "}
          Extension!
        </h2>
        <input onChange={(e) => setData(e.target.value)} value={data} />
        <br/>
        <Button type="primary">点我应该没反应才对</Button>
        <br/>
        <button onClick={() => {
          chrome.tabs.create({
            url: './tabs/delta-flyer.html'
          })
        }}>open tab page</button>
        <br/>
        <button onClick={() => {
          chrome.tabs.query({active: true, currentWindow: true}, function(tabs) {
            const { id } = tabs[0]
            chrome.scripting.executeScript({
              target: { tabId: id },
              func: () => {
                const iframe = document.createElement("iframe")
                iframe.src = chrome.runtime.getURL("/tabs/delta-flyer.html")
                iframe.name = "delta-flyer"
                document.body.appendChild(iframe)
              }
            })
          })
        }}>iframe mounting</button>

      <br/>
      <button
        onClick={() => {
          iframeRef.current.contentWindow.postMessage("10 + 20", "*")
        }}>
        Trigger iframe eval
      </button>
      <iframe src="sandbox.html" ref={iframeRef} style={{ display: "none" }} />
    </div>
  </ThemeProvider>
  )
}

export default IndexPopup
