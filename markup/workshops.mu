div p-0
  {{ template "navbar" . }}
  {{ $verb := "register" }}
  {{ $guid := .guid }}
  {{ $emap := .emap }}
  {{ if .user }}
  {{ $verb = "grab" }}
  {{ end }}
  div flex flex-col md:flex-row space-x-9 items-start justify-center
    div w-full md:w-1/2
      div bg-yellow-600 text-black py-3 rounded-lg text-center text-2xl font-familjen my-3
        Workshops
      div mt-9
        div flex justify-between items-center
          div
            div flex space-x-3 items-center
              select select select-primary
                {{ range $i, $item := .options }}
                option
                  {{ $item }} 
                {{ end }}
              div
                August 18, 2024
          div
            4 hour event
        div whitespace-nowrap flex justify-center flex-wrap space-x-3 space-y-3 
          {{ range $i, $item := .nrgs }}
          {{ $disabled := "disabled" }}
          {{ if index $emap $item.Id }}
            {{ $disabled = "outline" }}
          {{ end }}
          {{ if eq $i 0 }}
            div ml-3 mt-3
              a w-32 btn btn-{{$disabled}} href=/{{$verb}}/{{$guid}}/{{$item.Id}}
                {{ $item.Label }}
          {{ else }}
            div 
              a w-32 btn btn-{{$disabled}} href=/{{$verb}}/{{$guid}}/{{$item.Id}}
                {{ $item.Label }}
          {{ end }}
          {{ end }}
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
