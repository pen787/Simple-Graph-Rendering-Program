-- Any code outside this will be run when script is load
print("Quadratoc Script load!")
-- Run when graph load
function load()
  return {R=100,G=0,B=255}
end

-- Run for each x coordination
function render(x)
  return x * x
end
