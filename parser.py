import time

import psycopg2
from selenium import webdriver
from selenium.webdriver.common.by import By

drv = webdriver.Chrome()
conn = psycopg2.connect("dbname='lila' user='postgres' host='localhost' password='kolyan2811'")
cur = conn.cursor()

with open("res_hse.txt", encoding="utf8") as line:
    arr = [i.split("\n") for i in line.read().split("\n\n\n")]
    for i in arr:
        cur.execute("INSERT INTO univ (un, way) VALUES (%s, %s)", ("HSE", i[0]))
        for j in i[1:]:
            cur.execute("INSERT INTO ways (way, subjects) VALUES (%s, %s)", (i[0], ', '.join(i[1:])))
            drv.get(f"https://lifehacker.ru/kursy/courses?search={j}")
            courses = []

            for k in range(5):
                try:
                    element = drv.find_elements(By.CLASS_NAME, "b-btn--secondary")[k].get_attribute("href")
                except Exception:
                    drv.get(f"https://lifehacker.ru/kursy/courses?search={j}")
                    continue
                drv.get(element)

                title = drv.find_element(By.CLASS_NAME, "l-course__title").text
                feedback = drv.find_element(By.CLASS_NAME, "l-course__feedback").text.split("\n")
                feedback.append("-")
                features = drv.find_element(By.CLASS_NAME, "l-features").text.split("\n")
                price = drv.find_element(By.CLASS_NAME, "item-price").text
                link = drv.find_element(By.CLASS_NAME, "l-course__button").get_attribute("href")
                try:
                    skills = list(filter(lambda x: x not in "1234567890", drv.find_element(By.CLASS_NAME, "l-skills__items").text.split("\n")))
                except Exception:
                    skills = []

                try:
                    cur.execute("INSERT INTO courses (name, rating, description, difficulty, duration, skills, price, rating_number) VALUES (%s, %s, %s, %s, %s, %s, %s, %s)",
                            (title, feedback[0], link, features[1], feedback[2], ', '.join(skills), price, feedback[1]))
                except Exception:
                    drv.get(f"https://lifehacker.ru/kursy/courses?search={j}")
                    continue
                drv.get(f"https://lifehacker.ru/kursy/courses?search={j}")

            cur.execute("INSERT INTO subjects (name, courses) VALUES (%s, %s)", (j, ', '.join(courses)))
            conn.commit()
cur.close()
conn.close()
