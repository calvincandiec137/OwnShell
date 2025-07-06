#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/wait.h>
#include <unistd.h>

char *cwd;

void shell_loop();
char *read_line();
char **parse_line(char *line);
int handle_builtin(char **args);
int execute_commands(char **args);

int main() {
  cwd = malloc(100);
  shell_loop();
  free(cwd);
  return 0;
}

void shell_loop() {
  while (1) {
    getcwd(cwd, 100);
    printf("msh%s> ", cwd);

    char *command = read_line();
    char **commands = parse_line(command);

    if (commands[0] == NULL)
      continue;

    if (strcmp(commands[0], "cd") == 0) {
      handle_builtin(commands);
    } else {
      execute_commands(commands);
    }

    free(command);
    free(commands);
  }
}

char *read_line() {
  char *command = malloc(100);
  fgets(command, 100, stdin);
  command[strcspn(command, "\n")] = '\0';
  return command;
}

char **parse_line(char *line) {
  char **commands = malloc(100 * sizeof(char *));
  int i = 0;
  char *token = strtok(line, " ");
  while (token) {
    commands[i++] = token;
    token = strtok(NULL, " ");
  }
  commands[i] = NULL;
  return commands;
}

int handle_builtin(char **args) {
  if (args[1] == NULL) {
    fprintf(stderr, "cd: expected argument\n");
    return 1;
  }
  if (chdir(args[1]) != 0) {
    perror("test");
    return 1;
  }
  return 0;
}

int execute_commands(char **args) {
  pid_t pid = fork();

  if (pid == 0) {
    execvp(args[0], args);
  } else if (pid > 0) {
    wait(NULL);
  } else {
    perror("Fork issue.");
  }
  return 0;
}