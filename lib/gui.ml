open Tk

(* Constants *)
let application_name = "Chippotto"

let window_size = "640x320"

let screen_pixels_width = 64

let screen_pixels_height = 32

let screen_size = screen_pixels_width * screen_pixels_height

let pixel_size = 10

let primary_color = `Black

let secondary_color = `White

(* Window creation *)

let top = openTk () ;;

Wm.title_set top application_name ;;

Wm.geometry_set top window_size ;;

(* Canvas creation *)

let canvas =
  Canvas.create top ~background:`Black
    ~width:(screen_pixels_width * pixel_size)
    ~height:(screen_pixels_height * pixel_size)
in
pack [canvas] ~fill:`None ~expand:true

(* Events handling *)

let handle_mouse_event ev =
  match ev with `Motion -> print_string "a" |> ignore | _ -> ()

let start_mainloop () = Printexc.print mainLoop ()
