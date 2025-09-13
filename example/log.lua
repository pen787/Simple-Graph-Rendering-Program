-- Any code outside this will be run when script is load
print("Log Script load!")
-- Run when graph load
function load() end

-- Run for each x cords
function render(x)
  return math.log(x, 10)
end