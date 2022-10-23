package snapshot_test

import (
	"errors"
	"os"
	"testing"

	"github.com/kodiiing/snapshot"
)

func TestMatchSnapshot(t *testing.T) {
	t.Run("WithoutConfiguration", func(t *testing.T) {
		t.Cleanup(func() {
			err := os.Remove("test-without-configuration.snap")
			if err != nil {
				t.Logf("[cleanup] removing file: %s", err.Error())
			}
		})

		input := `it('will check the values and pass', () => {
const user = {
	createdAt: new Date(),
	name: 'Bond... James Bond',
};

expect(user).toMatchSnapshot({
	createdAt: expect.any(Date),
	name: 'Bond... James Bond',
});
});

// Snapshot
exports['will check the values and pass 1'] =
Object {
"createdAt": Any<Date>,
"name": 'Bond... James Bond',
}
;`
		ok, err := snapshot.MatchSnapshot("test-without-configuration", input, snapshot.Config{})
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())
		}

		if !ok {
			t.Error("expected `ok` to be true, got false")
		}
	})

	t.Run("ExistingFile", func(t *testing.T) {
		t.Cleanup(func() {
			err := os.Remove("test-existing-file-1.snap")
			if err != nil {
				t.Logf("[cleanup] removing file: %s", err.Error())
			}
		})

		loremIpsum := `Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur iusto beatae nobis, 
eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum tenetur
	ut sapiente sequi culpa ea. Consequatur maxime quae commodi. Fuga asperiores eveniet 
	dicta voluptas magni exercitationem officia maxime possimus nam rerum laboriosam nihil 
	officiis saepe, dolor cupiditate nesciunt maiores qui. Illum facere at rerum vitae.
Voluptatem aliquam fugit maiores ex repellendus provident dolores nesciunt voluptas
dicta, doloribus magni laudantium inventore vitae. Tenetur est sunt doloribus quod culpa
	error alias sint in magni mollitia. Numquam, ex. Quod nobis ut tenetur pariatur atque,
	quis inventore eius quaerat debitis odit molestias ipsa perspiciatis commodi praesentium
	explicabo vel modi nesciunt! Veniam nobis recusandae quidem quod expedita sit eaque! Esse?
Magnam recusandae vero ut omnis ipsa esse, neque mollitia aliquid tempora corrupti nam 
officia itaque reiciendis sit obcaecati dicta ducimus doloribus ad molestias culpa? 
Molestiae vero repudiandae quidem inventore fugiat? Explicabo exercitationem fugit quo 
tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis autem
pariatur non natus, amet nam similique omnis et cumque beatae consectetur expedita ab unde!
Totam temporibus neque, fugiat voluptatem unde facilis enim explicabo nihil illum nesciunt 
velit amet ducimus soluta? Blanditiis distinctio cum hic aspernatur, voluptas officia 
eligendi iste tempore sit dolores obcaecati magnam.`

		err := os.WriteFile("test-existing-file-1.snap", []byte(loremIpsum), 0755)
		if err != nil {
			t.Fatalf("writing initial file: %s", err.Error())
		}

		ok, err := snapshot.MatchSnapshot("test-existing-file-1", loremIpsum, snapshot.Config{})
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())

			var snapshotError snapshot.SnapshotError
			if errors.As(err, &snapshotError) {
				t.Errorf("expecting: %s", snapshotError.Snapshot)
				t.Errorf("got: %s", snapshotError.Received)
			}
		}

		if !ok {
			t.Error("expected `ok` to be true, got false")
		}
	})

	t.Run("AlwaysUpdate", func(t *testing.T) {
		t.Cleanup(func() {
			err := os.Remove("test-always-update-1.snap")
			if err != nil {
				t.Logf("[cleanup] removing file: %s", err.Error())
			}
		})

		loremIpsum := `Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur iusto beatae nobis, 
eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum tenetur
	ut sapiente sequi culpa ea. Consequatur maxime quae commodi. Fuga asperiores eveniet 
	dicta voluptas magni exercitationem officia maxime possimus nam rerum laboriosam nihil 
	officiis saepe, dolor cupiditate nesciunt maiores qui. Illum facere at rerum vitae.
Voluptatem aliquam fugit maiores ex repellendus provident dolores nesciunt voluptas
dicta, doloribus magni laudantium inventore vitae. Tenetur est sunt doloribus quod culpa
	error alias sint in magni mollitia. Numquam, ex. Quod nobis ut tenetur pariatur atque,
	quis inventore eius quaerat debitis odit molestias ipsa perspiciatis commodi praesentium
	explicabo vel modi nesciunt! Veniam nobis recusandae quidem quod expedita sit eaque! Esse?
Magnam recusandae vero ut omnis ipsa esse, neque mollitia aliquid tempora corrupti nam 
officia itaque reiciendis sit obcaecati dicta ducimus doloribus ad molestias culpa? 
Molestiae vero repudiandae quidem inventore fugiat? Explicabo exercitationem fugit quo 
tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis autem
pariatur non natus, amet nam similique omnis et cumque beatae consectetur expedita ab unde!
Totam temporibus neque, fugiat voluptatem unde facilis enim explicabo nihil illum nesciunt 
velit amet ducimus soluta? Blanditiis distinctio cum hic aspernatur, voluptas officia 
eligendi iste tempore sit dolores obcaecati magnam.`

		err := os.WriteFile("test-always-update-1.snap", []byte(loremIpsum), 0755)
		if err != nil {
			t.Fatalf("writing initial file: %s", err.Error())
		}

		ok, err := snapshot.MatchSnapshot("test-always-update-1", loremIpsum, snapshot.Config{AlwaysUpdate: true})
		if err != nil {
			t.Errorf("unexpected error: %s", err.Error())

			var snapshotError snapshot.SnapshotError
			if errors.As(err, &snapshotError) {
				t.Errorf("expecting: %s", snapshotError.Snapshot)
				t.Errorf("got: %s", snapshotError.Received)
			}
		}

		if !ok {
			t.Error("expected `ok` to be true, got false")
		}
	})

	t.Run("InvalidPermission", func(t *testing.T) {
		t.Cleanup(func() {
			err := os.Remove("test-invalid-permission.snap")
			if err != nil {
				t.Logf("[cleanup] removing file: %s", err.Error())
			}

		})

		err := os.WriteFile("test-invalid-permission.snap", []byte("Nothing to see here"), 0000)
		if err != nil {
			t.Fatalf("creating initial file: %s", err.Error())
		}

		ok, err := snapshot.MatchSnapshot("test-invalid-permission", "Nothing to see here", snapshot.Config{})
		if err == nil {
			t.Errorf("expecting an error, got nil instead: %s", err.Error())
		}

		if !errors.Is(err, os.ErrPermission) {
			t.Errorf("expecting error of `os.ErrPermission`, instead got %s", err.Error())
		}

		if ok {
			t.Error("expecting ok to be false, got true")
		}
	})

	t.Run("Difference", func(t *testing.T) {
		t.Cleanup(func() {
			err := os.Remove("test-difference-1.snap")
			if err != nil {
				t.Logf("[cleanup] removing file: %s", err.Error())
			}
		})

		correctLoremIpsum := `Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur iusto beatae nobis, 
eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum tenetur
	ut sapiente sequi culpa ea. Consequatur maxime quae commodi. Fuga asperiores eveniet 
	dicta voluptas magni exercitationem officia maxime possimus nam rerum laboriosam nihil 
	officiis saepe, dolor cupiditate nesciunt maiores qui. Illum facere at rerum vitae.
Voluptatem aliquam fugit maiores ex repellendus provident dolores nesciunt voluptas
dicta, doloribus magni laudantium inventore vitae. Tenetur est sunt doloribus quod culpa
	error alias sint in magni mollitia. Numquam, ex. Quod nobis ut tenetur pariatur atque,
	quis inventore eius quaerat debitis odit molestias ipsa perspiciatis commodi praesentium
	explicabo vel modi nesciunt! Veniam nobis recusandae quidem quod expedita sit eaque! Esse?
Magnam recusandae vero ut omnis ipsa esse, neque mollitia aliquid tempora corrupti nam 
officia itaque reiciendis sit obcaecati dicta ducimus doloribus ad molestias culpa? 
Molestiae vero repudiandae quidem inventore fugiat? Explicabo exercitationem fugit quo 
tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis autem
pariatur non natus, amet nam similique omnis et cumque beatae consectetur expedita ab unde!
Totam temporibus neque, fugiat voluptatem unde facilis enim explicabo nihil illum nesciunt 
velit amet ducimus soluta? Blanditiis distinctio cum hic aspernatur, voluptas officia 
eligendi iste tempore sit dolores obcaecati magnam.`

		wrongLoremIpsum := `Lorem ipsum dolor sit amet consectetur adipisicing elit. Consequatur iusto beatae nobis, 
eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum
	ut sapiente sequi culpa ea. Consequatur maxime quae commodi. Fuga asperiores eveniet 
	dicta voluptas magni exercitationem officia maxime possimus nam rerum laboriosam nihil 
	officiis saepe, dolor cupiditate nesciunt maiores qui. Illum facere at rerum vitae.
Voluptatem aliquam fugit maiores ex repellendus provident dolores nesciunt voluptas
dicta, doloribus magni laudantium inventore vitae. Tenetur est sunt doloribus quod culpa
	error alias sint in magni mollitia. Numquam, ex. Quod nobis ut tenetur pariatur atque,
	quis inventore eius quaerat debitis odit molestias ipsa perspiciatis commodi praesentium
	explicabo vel modi nesciunt! Veniam nobis recusandae quidem quod expedita sit eaque! Esse?
Magnam recusandae vero ut omnis ipsa esse, neque mollitia aliquid tempora corrupti nam 
officia itaque reiciendis sit obcaecati dicta ducimus doloribus ad molestias culpa? 
Molestiae vero repudiandae quidem inventore fugiat? Explicabo exercitationem fugit quo 
tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis
pariatur non natus, amet nam similique omnis et cumque beatae consectetur expedita ab unde!
Totam temporibus neque, fugiat voluptatem unde facilis enim explicabo nihil illum nesciunt 
velit amet ducimus soluta? Blanditiis distinctio cum hic aspernatur, voluptas officia 
eligendi iste tempore sit dolores obcaecati magnam.`

		err := os.WriteFile("test-difference-1.snap", []byte(correctLoremIpsum), 0755)
		if err != nil {
			t.Fatalf("writing initial file: %s", err.Error())
		}

		ok, err := snapshot.MatchSnapshot("test-difference-1", wrongLoremIpsum, snapshot.Config{})
		if err == nil {
			t.Errorf("expecting an error, got nil instead")
		}

		var snapshotError snapshot.SnapshotError
		if errors.As(err, &snapshotError) {
			if snapshotError.Difference != 2 {
				t.Errorf("expecting `snapshot.Difference` to be 2, got %d instead", snapshotError.Difference)
			}

			expectedSnapshot := `2 | eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum tenetur
...
14 | tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis autem`
			if snapshotError.Snapshot != expectedSnapshot {
				t.Errorf("expecting `snapshotError.Snapshot` to match %s, got %s", expectedSnapshot, snapshotError.Snapshot)
			}

			expectedReceived := `2 | eaque tempora eius amet repellat, ut dolor voluptate laudantium fuga quis. Harum
...
14 | tempore dolores adipisci, iste qui doloribus ut ducimus a reprehenderit. Reiciendis`
			if snapshotError.Received != expectedReceived {
				t.Errorf("expecting `snapshotError.Received` to match %s, got %s", expectedReceived, snapshotError.Received)
			}
		}

		if ok {
			t.Error("expected `ok` to be false, got true")
		}
	})
}
