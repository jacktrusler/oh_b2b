# Must have tmux running, splits into 3 windows and runs watchers for hot reloading
# Generates templ files to go and tailwind to output on the fly
tmux-dev: 
  tmux split-window -h \; \
    split-window -v \; \
    select-pane -t 0 \; send-keys 'air' C-m \; \
    select-pane -t 1 \; send-keys 'templ generate --watch' C-m \; \
    select-pane -t 2 \; send-keys 'tailwindcss -i public/input.css -o public/output.css --watch' C-m \; \
