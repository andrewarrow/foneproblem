div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full ml-16 mr-16
      div text-center text-2xl font-familjen my-3 bg-red-800 rounded-lg
        do you have a foneproblem?
      div
        a href=/core/start
          div flex flex-wrap items-center space-x-4 space-y-4
            div w-64
              img src=https://i.imgur.com/jmvTY2t.png rounded-full
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
      p text-center mt-16 mb-64
        <a class="btn btn-primary" href="/core/start">Enter</a>
