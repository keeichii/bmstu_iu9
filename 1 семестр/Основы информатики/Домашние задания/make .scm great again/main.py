import os

def save_file(file_path: str, content: str, anti_plag: str, teacher_comment: str) -> bool:
    """Сохраняет содержимое в указанный файл и проверяет, успешно ли это выполнено."""
    try:
        # Проверяем, существует ли директория
        dir_path = os.path.dirname(file_path)
        if not os.path.isdir(dir_path):
            print(f'\'{dir_path}\' не является директорией. Пожалуйста, укажите корректный путь к директории.')
            return False

        with open(file_path, 'w') as file:
            file.write(content)  # Записываем содержимое в файл
            # Добавляем комментарии в конец файла
            file.write(f"\n;{anti_plag}\n;{teacher_comment}\n")

        return os.path.exists(file_path) and os.path.getsize(file_path) > 0
    except Exception as e:
        print(f"Произошла ошибка: {e}")
        return False

def main():
    # Запрашиваем путь сохранения директории
    dir_path = input("Введите путь к директории для сохранения файлов (или 'stop' для выхода): ")
    if dir_path.lower() == "stop":
        return  # Завершение программы

    while True:
        print("Введите содержимое файла (первая строка - название файла, затем антиплагиат, комментарий и код).")
        print("Введите 'over' для завершения ввода:")

        lines = []
        while True:
            line = input()  # Чтение строки
            if line.strip().lower() == "over":  # Если команда "over"
                break
            lines.append(line)

        # Проверка на количество строк
        if len(lines) < 5:
            print("Недостаточно строк для файла. Пожалуйста, введите как минимум 5 строк.")
            continue

        # Получаем название файла из первой строки
        file_name = lines[0].strip() + ".scm"
        file_path = os.path.join(dir_path, file_name)

        # Ищем нужные строки по ключевым словам
        anti_plag = ""
        teacher_comment = ""
        code_lines = []

        for line in lines:
            if "антиплагиат" in line.lower():
                anti_plag = line
            elif "комментарий" in line.lower():
                teacher_comment += line + "\n"
            elif "код решения" not in line.lower():
                code_lines.append(line)

        if anti_plag == "" or teacher_comment == "":
            print("Не найдены необходимые строки для антипластиата или комментария преподавателя.")
            continue

        # Удаляем начальные пустые строки из code_lines
        while code_lines and code_lines[0].strip() == "":
            code_lines.pop(0)

        # Код начинается с 5-й строки
        code = "\n".join(code_lines)  # Собираем код

        # Сохраняем файл и проверяем результат
        success = save_file(file_path, code, anti_plag, teacher_comment.strip())
        print("Сохранение:", success)

if __name__ == "__main__":
    main()
