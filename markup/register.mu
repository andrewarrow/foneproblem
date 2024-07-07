div p-0 
  {{ template "navbar" . }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div bg-yellow-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        Register
      div mt-9
        div flex justify-between
          div
            West LA August, 18 2024
          div
            4 hour event
        div flex justify-center flex-wrap space-x-3 space-y-3 
          div ml-3 mt-3
            a w-32  btn btn-selected href=/
              Female energy 18-24
        div flex justify-center flex-wrap space-x-3 space-y-3 
          div mt-3 text-yellow-500
            There is a spot for you!
      div mt-9 flex justify-center
        form mt-6 space-y-3 id=register
          div
            input value={{.email}} input input-primary id=email placeholder=email autofocus=true
          div
            input value=testing123 input input-primary id=password placeholder=password
          div flex justify-center
            input type=submit btn btn-primary value=Save-My-Spot
