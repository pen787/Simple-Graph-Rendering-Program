-- Any code outside this will be run when script is load
print("Randomize Linear Script load!")

local randomNumber = 0
local randomK = 0
-- Run when graph load
function load()
  randomNumber = math.random(-10, 10)
  randomK = math.random(-10, 10)
end

-- Run for each x cords
function render(x)
  return (randomK*x) + randomNumber
end