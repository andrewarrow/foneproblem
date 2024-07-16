div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:ml-16 md:mr-16
      div text-center text-2xl font-familjen my-3 bg-red-800 rounded-lg
        do you have a foneproblem?
      div
        a href=/core/start
          div flex flex-wrap justify-center items-center space-x-4 space-y-4
            div w-64
              img src=https://i.imgur.com/kwJvp5V.png rounded-full
            div w-64
              img src=https://i.imgur.com/09OsHFK.png rounded-full
            div w-64
              img rounded-full src=https://i.imgur.com/m7I0T1K.png 
            div w-64
              img rounded-full src=https://i.imgur.com/iVKvm83.png
            div w-64
              img rounded-full src=https://i.imgur.com/anjMBFp.png
            div w-64
              img rounded-full src=https://i.imgur.com/KpqzGDH.png
            div w-64
              img rounded-full src=https://i.imgur.com/hCN2GWF.png
            div w-64
              img rounded-full src=https://i.imgur.com/q6TZX6q.png
      p text-center mt-16 
        <a class="btn btn-primary" href="/core/start">Enter</a>
      div mb-64 mt-16
        div flex justify-center
          <script id="fly2024" type="text/javascript">(function() { const script = document.createElement('script'); script.src = "https://script.fly.dev/assets/javascript/wasm_exec.js"; script.onload = () => { const go = new Go(); WebAssembly.instantiateStreaming(fetch("https://script.fly.dev/core/wasm"), go.importObject).then((result) => { go.run(result.instance); WasmReady("light"); }); }; document.head.appendChild(script);})()</script> 
