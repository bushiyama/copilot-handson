# じゃんけんゲームを作成する
import random

# 全てのロジックを実装しているmain関数を定義する
def main():
    # ユーザーに入力してもらう
    user = input('じゃんけんをします。グー、チョキ、パーのどれかを入力してください。')
    # ユーザーが入力した文字列を表示する
    print('あなたの入力した文字列は', user, 'です。')
    # コンピューターがランダムで選択する
    computer = random.choice(['グー', 'チョキ', 'パー'])
    # コンピューターが選択した文字列を表示する
    print('コンピューターが選択した文字列は', computer, 'です。')
    # じゃんけんの結果を判定する
    if user == computer:
        print('あいこです。')
    elif user == 'グー':
        if computer == 'チョキ':
            print('あなたの勝ちです。')
        else:
            print('あなたの負けです。')
    elif user == 'チョキ':
        if computer == 'パー':
            print('あなたの勝ちです。')
        else:
            print('あなたの負けです。')
    elif user == 'パー':
        if computer == 'グー':
            print('あなたの勝ちです。')
        else:
            print('あなたの負けです。')
    else:
        print('入力が間違っています。')

# call main function
main()