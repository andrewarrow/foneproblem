{{ define "navbar" }}
  div navbar bg-base-200 font-familjen
    div navbar-start 
      div btn btn-ghost text-4xl
        a href=/
          img src=logo.png w-12
    div navbar-center flex hidden md:block
    div navbar-end
      div flex space-x-3
        {{ if .user }}
          a href=/core/dashboard
            Dashboard
          a href=/ id=logout
            Logout
        {{ else }}
          a href=/core/workshops link link-hover
            Register
        {{ end }}
  {{ end }}
