package main

import "fmt"

func main() {
  /* 
  
  Es ist 1996 und du fragst dich, 
  "Soll ich's wirklich machen oder laß ich's lieber sein? Jein ..." 
 
  - Fettes Brot, Jein -
  */

  Intro()

  Jein(Solist("Die kleinen Schweinerein"))
  Jein(Solist("Mit der Freundin des besten Freunds"))
  Jein(Duett("Lass diese Party sein!"))
  Jein("Schlaghosen und Buffalos")

  End()

}

func Jein(jein interface{}) {
  if tun, ok := jein.(LassEsUnsTun); ok {
    tun.Schweinerein()
    return
  }

  switch t := jein.(type) {
    default:
      MannMitBart(t.(string), "Kann man machen oder halt auch nicht")
    case Solist:
      MannMitBart(string(t), "Wenn ich du wäre dann wäre ich lieber ich")
  }
}

type Solist string
type Duett string
type LassEsUnsTun interface {
  Schweinerein()
}

func (_ Duett) Schweinerein() {
  fmt.Println("    > Schweinerrein?")
  fmt.Println("    < *räusper*\n")
}

func MannMitBart(ask, answ string) {
  // Weise ist der Mann mit langem Bart!
  fmt.Println("    >", ask, "?")
  fmt.Println("    <", answ, "!\n")
}

func Intro() {
  fmt.Println(`
    Jein

    Es ist 1996,
    meine Freundin ist weg und bräunt sich 
    in der Südsee. Allein?
    Ja, mein Budget war klein.
    Na fein! Herei, willkommen im Verein!
    Ich wette, heute machen wir erneute fette Beute
    treffen seute Bräute und lauter nette Leute.
    Warum dauernd trauern?
    Wow, schaut euch diese Frau an!
    Schande, dazu bist du imstande?!
    Kaum ist deine Herzallerliebste aus dem Lande
    und du Hengst denkst längst an ‘ne andre.
    Soll ich denn heulen?
    Ihr wißt, daß ich meiner Freundin treu bin.
    Ich bin brav aber ich traf eben my first love.
    Ich darf zwar nur im Schlaf, 
    doch auf sie war ich schon immer scharf.
    Habt ihr den Blick geahnt, 
    den sie mir eben durchs Zimmer warf!
    Oh, mein Gott, wat hat der Trottel Sott.
    What a Pretty Woman, 
    das Glück is’ mit die Dummen.
    Wenn ich die stummen Blicke schicke,
    sie wie Rummenigge kicke, meint ihr checkt sie das ?
    Du bist durchschaubar wie Plexiglas!
    Uh, sie kommt auf dich zu.
    Na Kleiner, hast du Bock auf Schweinereien
    Ja klar, äh nein, ich mein Jein !

    Soll ich´s wirklich machen oder laß ich´s lieber sein? Jein....

    Ich habe einen Freund - Ein guter ? - sozusagen, mein bester,
    und ich habe ein Problem, 
    ich steh auf seine Freundin - Nicht auf seine Schwester?
    Würd ich auf die Schwester stehen, hätt’ ich nicht das Problem
    das wir haben, wenn er, sie und ich uns sehen.
    Kommt sie in den Raum, wird mir schwindelig.
    Sag ich, sie will nichts von mir, dann schwindel ich.
    Ich will sie, sie will mich, das weiß sie, das weiß ich.
    Nur mein bester Freund, der weiß das nicht.
    Und somit sitz ich sozusagen in der Zwickmühle
    und das ist auch der Grund, 
    warum ich mich vom Schicksal gefickt fühle.
    Warum hat er die schönste Frau zur Frau?
    Mit dem schönsten Körperbau! - Und ist sie schlau? - Genau!
    Es steigen einem die Tränen in die Augen, wenn man sieht
    was mit mir passiert und was mit mir geschieht.
    Es erscheinen Engelchen und Teufelchen auf meiner Schulter
    Engel links, Teufel rechts:
    ‘Lechz, nimm dir die Frau, sie will es doch auch
    kannst du mir erklären, wozu man gute Freunde braucht?’
    ‘Halt, der will dich linken", schreit der Engel von der Linken,
    „weißt du nicht, daß sowas Scheiße ist und Lügner stinken?’
    Und so streiten sich die beiden um mein Gewissen.
    Und ob ihr’s glaubt oder nicht, mir geht es echt beschissen.
    Und während sich die beiden anschreien,
    entscheide ich mich für Ja, Nein, ich mein Jein !
    Soll ich´s wirklich machen oder laß ich´s lieber sein? Jein....

    Ich schätz, jetzt bin ich der Solist in unserem Knabenchor.
    - Ey, Schiff, was hast’ denn heute abend vor?
    Ich mach hier nur noch meine Strophe fertig,
    pack meine sieben Sachen und dann werd ich
    mich zu meiner Freundin begeben,
    denn wenn man ehrlich gesteht,
    sind solche netten, ruhigen Abende eher spärlich gesät
    - Aha, dabei biste eingeladen,
    auf das beste aller Feste auf der Gästeliste eingetragen!
    - Und wenn du nicht mitkommst dann hast
    du echt was verpaßt. Und wen wundert’s? 
    Es wird fast die Party des Jahrhunderts.
    Mmh, Lust hätt ich ja eigentlich schon!
    Oh, es klingelt just das Telefon.
    -Und sie sacht, es wär schön, 
    wenn du bei mir bleibst heut nacht
    ich dacht, das wär abgemacht?
    Wißt ihr, ich liebe diese Frau und deswegen
    komm ich von der Traufe in den Regen.
    -Na was ist nun Schiffmeister, 
    kommst du mit, du Kollegenschwein.
    Ja, äh Nein, ich mein Jein!

    Soll ich´s wirklich machen oder laß ich´s lieber sein? Jein....
    `)
}

func End() {
  fmt.Println("    - Fettes Brot, Jein -\n")
}
