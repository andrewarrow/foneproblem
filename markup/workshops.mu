div p-0 
  {{ template "navbar" . }}
  {{ $verb := "register" }}
  {{ if .user }}
  {{ $verb = "grab" }}
  {{ end }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div bg-yellow-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        Workshops
      div mt-9
        div flex justify-between
          div
            West LA August, 18 2024
          div
            4 hour event
        div flex justify-center flex-wrap space-x-3 space-y-3 
          div ml-3 mt-3
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/f18
              Female energy 18-24
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/m18
              Male energy 18-24
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/h18
              Human energy 18-24
          div ml-3 mt-3
            a w-32  btn btn-disabled href=/{{$verb}/{{.guid}}/f25
              Female energy 25-29
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/m25
              Male energy 25-29
          div 
            a w-32  btn btn-disabled href=/{{$verb}/{{.guid}}/h25
              Human energy 25-29
          div ml-3 mt-3
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/f30
              Female energy 30-39
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/m30
              Male energy 30-39
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/h30
              Human energy 30-39
          div ml-3 mt-3
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/f40
              Female energy 40-49
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/m40
              Male energy 40-49
          div 
            a w-32  btn btn-outline href=/{{$verb}/{{.guid}}/h40
              Human energy 40-49
        div mt-3 flex justify-between text-yellow-600
          div
            We have 2 of 12 spots reserved.
          div
            Get your spot today.
        div mt-6 flex justify-between 
          div
            We do groups of 12 with 4 female energy, 4 male energy, and 4 human
            energy. Anyone is welcome to select which engery they resinate with
            the most. But ages are ages, you have to select your correct age.
        div mt-6 flex justify-between 
          div
            If there is nothing left for your age group, move on to the 
            <a href="/" class="link">next event</a>
            and register early.
