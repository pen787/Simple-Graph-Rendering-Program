-- Any code outside this will be run when script is load
print("Quadratoc Script load!")
-- Run when graph load
function load()
  print("Hello!")
end

-- Run for each x cords
function render(x)
  return x * x
end