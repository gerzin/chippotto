(lang dune 3.10)

(name chippotto)

(generate_opam_files true)

(source
 (github gerzin/chippotto))

(authors "Gerardo Zinno")

(maintainers "Gerardo Zinno")

(license MIT)

(documentation https://url/to/documentation)

(package
 (name chippotto)
 (synopsis "A simple Chip8 interpreter in Ocaml")
 (description "A simple Chip8 interpreter in Ocaml")
 (depends ocaml dune labltk)
 (tags
  (topics "to describe" your project)))

; See the complete stanza docs at https://dune.readthedocs.io/en/stable/dune-files.html#dune-project
