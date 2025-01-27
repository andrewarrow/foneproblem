div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div bg-yellow-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        do you have a foneproblem?
      div bg-purple-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        do you even take your phone with you inside a SAUNA?
      div space-y-6 pt-9
        {{ range $i, $item := .questions }}
          p border-b-2 border-gray-700 pb-3
            {{add 1 $i}}. {{$item}}
        {{ end }}
      div text-center mt-9 pb-64
        a href=/core/workshops btn btn-primary
          Workshops In Your Area
