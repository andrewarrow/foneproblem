div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div text-center text-2xl font-familjen my-3
        do you have a foneproblem?
      div
        a href=/core/start
          img src=https://i.imgur.com/R0TJtnO.png rounded-full
      p text-center mt-6
        <a class="btn btn-primary" href="/core/start">Enter</a>
