#!/bin/bash

session="aoc"

tmux new-session -d -s $session

window=1
tmux rename-window -t $session:$window 'code'
tmux send-keys -t $session:$window 'cd ~/Desktop/aod && nvim .' C-m

window=2
tmux new-window -t $session:$window -n 'run'
tmux send-keys -t $session:$window 'cd ~/Desktop/aoc && go run main.go y15.go' C-m

tmux attach-session -t $session
