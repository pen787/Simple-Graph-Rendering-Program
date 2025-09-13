--Any code outside this will be run when script is load
print("Script load!")
-- Run when graph load
function load()
  print("hello!")
end

-- Run for each x cords
function render(x)
  return x*x
end