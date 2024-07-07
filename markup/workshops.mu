div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div bg-yellow-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        Workshops
      div mt-9
        div flex justify-between
          div
            West LA August 2024
          div
            4 hour event
        div flex justify-center flex-wrap space-x-3 space-y-3
          div ml-3 mt-3
            a w-32 h-64 btn-lg text-gray-600 btn btn-outline href=/roadmapp/register/3
              Female Energy 18-24
          div 
            a w-32 h-64 btn-lg text-gray-600 btn btn-outline href=/roadmapp/register/3
              Male Energy 18-24
          div 
            a w-32 h-64 btn-lg text-gray-600 btn btn-outline href=/roadmapp/register/3
              Human Energy 18-24
