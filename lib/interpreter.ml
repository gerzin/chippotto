(* Type definitions *)
type register = Value of int | Address of int

type instruction =
  | Ret
  | Clear
  | Jump of register
  | Call of register
  | Skip of register * int
  | Unknown

(* Utility to extract the subsections from an opcode*)
let extract code =
  let op = (code land 0xF000) asr 12 in
  let kk = code land 0x00FF in
  let x = (code land 0x0F00) asr 8 in
  let y = (code land 0x00F0) asr 4 in
  let n = code land 0x000F in
  let addr = code land 0x0FFF in
  (op, kk, x, y, n, addr)

let decode code =
  let op, kk, x, y, n, addr = extract code in
  match op with
  | 0x00 -> ( match code with 0x00E0 -> Clear | 0x00EE -> Ret | _ -> Unknown)
  | 0x01 -> Jump (Address addr)
  | 0x02 -> Call (Address addr)
  | 0x03 -> Skip (Value x, kk)
  | _ -> Unknown

(* ram memory *)
let ram = Array.make 0x1000 0x00

(*the v registers*)
let v_registers = Array.make 0x10 0x00

(* I register *)
let i = ref 0x00

(*program counter*)
let pc = ref 0x00

(*stack and stack pointer*)
let stack = Array.make 0x10 0x00
let sp = ref 0x200
